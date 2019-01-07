# API工具

公司内部的默认通信约定为：获取请求键值对中的```ActionID```，根据ActionID对应的函数，
Action函数应当能获取到当前请求上下文的信息，而一个应用中，不同路由下的请求上下文略有不同。

为了满足以上需求，dcapi设计了Action和Context接口，可根据不同需求，实现自己的Action和Context。

## Quick Start

## 获取包

```go get gitlab.dianchu.cc/goutil/dcapi.v2```

完整示例代码请移步dcapi.v2/_example

### 新建Manager

Manager用于管理Context要执行的handler和Action的映射

```go
AM = dcapi.New()
//Use指定了当前环境需要用到的handler，必须Use DoAction()才能执行Action
AM.Use(dcapi.Recovery(), dcapi.DoAction())
```

### 自定义Context

Context为当前请求的上下文，设计思路基于gin，通过自定义Context，可为Context结构体追加需求的成员和方法

```go
type GameContext interface {
    dcapi.Context

    Lang() string       //请求的必备参数
    RoleId() string     //请求的必备参数
}

type gameContext struct {
    dcapi.Context

    lang string
    roleId	string
}

func (s *baseCtx) RoleId() string{
    return s.roleId
}

func (s *baseCtx) Lang() string{
    return s.lang
}
```

### 编写通信协议的handler

不管是grpc服务还是http服务，均有受理请求的handler，在handler中，
我们需要获取当前请求对应的Action并构造Context。

以下为使用gin的示例

```go
func httpErr(c *gin.Context, info string){
	httpResp(c, 1, info, gin.H{})
}

func httpResp(c *gin.Context, code int, info string, data gin.H){
	c.JSON(http.StatusOK, gin.H{"Code": code, "Info": info, "Data": data})
}

func HttpHandler() gin.HandlerFunc{
	return func(c *gin.Context) {
		rawData, err := c.GetRawData()
		if err != nil{
			httpErr(c, err.Error())
			return
		}
		exp.Try(func() {
			any := json.Get(rawData)

			// 必备参数校验
			actionId := any.Get("ActionId").ToString()
			roleId := any.Get("RoleId").ToString()
			lang := any.Get("Lang").ToString()

			if actionId == "" || roleId == "" || lang == ""{
				httpErr(c, "missing params")
				return
			}


			act, ok := AM.GetAction(actionId)
			if !ok{
				httpErr(c, "can not find action")
			}

			//构造自定义的Context
			ctx, err := NewContext(AM.Handlers(), lang, roleId)
			if err != nil{
				httpErr(c, err.Error())
			}

			//将重要通用参数写入Context
			traceId := c.GetHeader(TRACE_ID)
			if traceId == "" {
				any.Get(TRACE_ID).ToString()
			}
			ctx.Store().Set(TRACE_ID, traceId)
			ctx.Store().Set(dcapi.ACT, act)

			params, err := act.Params(rawData)
			if err != nil{
				httpErr(c, "params cannot unmarshal")
				return
			}
			ctx.SetParams(params)

			//开始执行Handler链
			ctx.Next()

			resp, err := ctx.Resp()
			if err != nil{
				httpErr(c, err.Error())
			}

			httpResp(c, resp.Code(), resp.Info(), resp.Data())
		}, func(ex exp.Exception){
			httpErr(c, "ServerError")
		})
	}
}

```

### 自定义Action

dcapi内未提供功能完整的Action，需要根据需求实现Action接口

```go
//Action执行的函数，此处的ctx为自定的GameContext
type DoFunc func(ctx GameContext)

type Action struct {
    dcapi.BaseAct   //BaseAct提供了ID()方法

    fn		DoFunc
    params	reflect.Type
}

//Params 将bytes反序列化为结构体的方法，具体反序列化方式视需求而定（xml? json?）
func (act *Action) Params(b []byte) (interface{}, error){
    p := reflect.New(act.params).Interface()
    err := json.Unmarshal(b, p)
    if err != nil{
        return nil, err
    }
    return p, nil
}

//Do 执行了Action回调，此处将自定义的GameContext作为参数传给回调
func (act *Action) Do(ctx dcapi.Context){
    c := ctx.(GameContext)
    act.fn(c)
}
```

### 编写Action回调函数

```go
//请求参数
type LoginParams struct {
	Jwt        string
	RetailId   string

    IsPanic    int8
    IsFailed   int8
}

//返回数据结构
type LoginResp struct {
	SystemTime		int64
}

//Login必须满足上面定义的DoFunc type
func Login(ctx game.GameContext) {
    //通用的重要参数可放在context.Context的Value中
	traceId := ctx.Value(game.TRACE_ID).(string)

    //获取请求参数的值
	params := ctx.Params().(*LoginParams)

    //自定义请求上下文的方法
	roleId := ctx.RoleId()

    //因为Use了Recovery()，直接panic不会中断程序
	if params.IsPanic == 1{
		panic("test panic")
	}

    //可用Failed或Failed快速返回错误
	if params.IsFailed == 1{
		dcapi.FailedWithErr(ctx, errors.New("test error"))
		return
	}

    //构建返回结构体
	resp := LoginResp{}
	resp.SystemTime = time.Now().Unix()

    //可用Success快速返回结构体
	dcapi.Success(ctx, resp)
}


//注册Action，可编写函数封装下面两句
//注意，该函数所在的包必须在服务启动时被import，否则不会注册Action
func init() {
	act := NewAction("10000", LoginParams, Login)
	AM.Register(act)
}
```

### 完成

服务启动后，便可接收请求，本示例的通信格式为Json，
故本示例Login Action的请求应当是：

```json
{
	"ActionId": 10000,
	"RoleId": "123",
	"AppId": 1014,
	"ServerId": 4,
	"Lang": "zh",
	"JWT": "str",
	"Device":{
		"DeviceId": "dev",
		"DeviceType": "type",
		"DeviceOS": "os"
	},
	"RetailId": "9"
}
```

成功后的返回

```json
{
    "Code": 0,
    "Data": {
        "SystemTime": 1530078499
    },
    "Info": "Success"
}
```

## 说明

### 请求参数
Action接口提供了将请求bytes反序列化为结构体的方法Params，而Context中则包含了该Params，所以在上例的Login(ctx GameContext)中，可以直接通过
```params := ctx.Params().(*LoginParams)```
获取请求参数的结构体。

其实请求的params类似web中的表单，go有丰富的第三方库可以验证结构体的值，因此在example包中，
使用[govalidator](https://github.com/asaskevich/govalidator)包对params进行验证，
并自定义了Valid() handler，再DoAction前验证params。

### 返回数据

dcapi的Response接口如下：

```go
//Response 为当前上下文返回的结果，携带了状态编号和提示信息
type Response interface {
	Set(key string, value interface{})
	Get(key string) interface{}

	Code() int				//当前响应的状态
	SetCode(code int)

	Info() string			//响应的提示信息
	SetInfo(info string)

	SetError(err error)		//根据err快速设置info和code

	Data() map[string]interface{} //将Response格式化为键值对
}
```

上例Login中的LoginResp结构体，通过Success装换后，写入到Response的Data中，这样当在handler中，
可以通过ctx.Resp()获得Response，Response中包含了结果状态码code、提示信息info，
以及结果数据data。并构建自己的返回结果。

## 参考

gin:   https://github.com/gin-gonic/gin
dcapi: https://gitlab.dianchu.cc/DevOpsGroup/goutils

## TODO
1. 添加测试用例
2. 丰富中间件