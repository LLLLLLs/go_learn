/*
Author : Haoyuan Liu
Time   : 2018/6/25
*/
package dcapi

//Response 为当前上下文返回的结果，携带了状态编号和提示信息
type Response interface {
	Set(key string, value interface{})
	Get(key string) interface{}

	Code() int //当前响应的状态
	SetCode(code int)

	Info() string //响应的提示信息
	SetInfo(info string)

	SetError(err error) //根据err快速设置info和code

	Data() map[string]interface{} //将Response格式化为键值对
}

//Action的返回
type response struct {
	code int
	info string
	data map[string]interface{}
}

func (r *response) Code() int {
	return r.code
}

func (r *response) SetCode(code int) {
	r.code = code
}

func (r *response) Info() string {
	return r.info
}

func (r *response) SetInfo(info string) {
	r.info = info
}

func (r *response) Set(key string, value interface{}) {
	r.data[key] = value
}

func (r *response) Get(key string) interface{} {
	v, ok := r.data[key]
	if !ok {
		return nil
	}
	return v
}

func (r *response) SetError(err error) {
	r.code = 1
	r.info = err.Error()
}

func (r *response) Data() map[string]interface{} {
	return r.data
}

func NewResponse() Response {
	return &response{
		0,
		"Success",
		make(map[string]interface{}),
	}
}
