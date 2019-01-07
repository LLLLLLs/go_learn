# 系统日志聚合

系统日志聚合服务由syslog服务器(rsyslog)收集程序运行时生成的日志，再转发给Fluentd后，经过相应处理，再次转发给Elasticsearch存储起来，然后在Kibana上对Elasticsearch中的数据进行搜索、查询、分析等操作。

## 日志传递

__日志输出流程图__：  

> 日志数据通路①：控制台输出  
> 日志数据通路②：Docker日志驱动经Fluentd输出（保留）  
> 日志数据通路③：TCP模式经syslog输出（优先）  

![日志输出流程图](https://gitlab.dianchu.cc/huangfulin/image_library/raw/master/system_log/access_doc/log_out_flowchart.png)  


## log_out_sdk_go简介  
`log_out_sdk_go`是一个日志输出的SDK，使用Go语言编写，并且也只可提供给Go项目调用。该SDK可以输出Go项目程序运行过程中生成的包含message、level_name、module、tag、trace_id等属性的日志到Fluentd并被转发到Elasticsearch中。当你在Kibana中查询时，每条日志的以上这些属性都会被单独列为Kibana中的Field（域），且在table中以表项形式显示。

## 安装log_output_go  

最新的日志输出SDK版本为`v1.1.0`。  

### 直接复制方式

从https://gitlab.dianchu.cc/chaos_go_sdk/log_out_sdk_go 下载最新的代码，把`log_out_sdk_go`放在项目相应目录下即可。  

### govendor安装方式

进入项目目录，打开终端。  

之前接入过该SDK的使用下面命令： 

```shell
$ go get gitlab.dianchu.cc/chaos_go_sdk/log_out_sdk_go 

$ govendor update gitlab.dianchu.cc/chaos_go_sdk/log_out_sdk_go 
```

未接入过的使用下面命令：

```shell
$ go get gitlab.dianchu.cc/chaos_go_sdk/log_out_sdk_go 

$ govendor add gitlab.dianchu.cc/chaos_go_sdk/log_out_sdk_go 
```

##  配置初始化

请使用`log_out_sdk_go`模块的初始化函数`InitLogger`进行初始化。使用容器部署时，还可以通过环境变量修改并覆盖该函数传入的配置参数（参见第3部分容器部署相关内容）。

## syslog需要在/data下创建syslog_buffer文件夹，请确保程序具备在/data下创建文件夹的权限，或者手动将文件夹的权限设置为666

```go
func (logger *CustomLogger) InitLogger(toStd bool, toSyslog bool, logLevel string, syslogIP string, syslogPort string, logMode string, loggerName string) (err error) 
```

### 初始化函数参数列表

| 参数名称   | 参数说明                                                     | 参数类型 | 初始值 |
| ---------- | ------------------------------------------------------------ | -------- | ------ |
| toStd      | 是否开启控制台打印，true打开，false关闭。                    | bool     | true   |
| toSyslog   | 是否开启syslog日志传输方式（可在Kibana中查询），true打开，false关闭。 | bool     | false  |
| logLevel   | 全局日志输出级别，同时应用于syslog日志和标准控制台输出。     | string   | INFO   |
| syslogIP   | syslog服务器TCP地址，必须填写，非容器部署时需要使用此IP。不同环境的syslog地址不同。 | string   | ""     |
| syslogPort | syslog服务器TCP端口，必须填写，非容器部署时需要使用此端口。一般请设置为514。 | string   | ""     |
| logMode    | 发往syslog服务器的日志输出模式，覆盖容器环境变量LOG_MODE，除非必要，一般请设置为""。 | string   | ""     |
| loggerName | 系统日志标签，需要填写成自己服务的名称，如dq2_chat。日志中会将其值赋给@fluentd_tag字段，用于区别不同服务。 | string   | ""     |


## 调用代码  

`example.go`：

```go
package main

import (
cLog "gitlab.dianchu.cc/chaos_go_sdk/log_out_sdk_go"
	"fmt"
    "time"
)

func main()  {
	cLog.Logger.InitLogger(true,true,"DEBUG","192.168.5.111","514","","fluentd_test")  // 全局必须且仅需进行一次初始化
	cLog.Info(&cLog.LogRecord{
		Message: fmt.Sprintln("This is a test information."),
	})
    time.Sleep(3 * time.Second) // 日志发送使用goroutine（协程），如果只发一条主进程就退出，需要sleep一下
}
```

##  日志输出函数

`cLog`为包`gitlab.dianchu.cc/chaos_go_sdk/log_out_sdk_go`的别名，可以任意指定。

```go
cLog.Logger.InitLogger(toStd bool, toSyslog bool, logLevel string, syslogIP string, syslogPort string, logMode string, loggerName string)  // 全局必须且仅需进行一次初始化

cLog.Debug(logRecord *LogRecord)     // DEBUG级别日志
cLog.Info(logRecord *LogRecord)      // INFO级别日志
cLog.Warning(logRecord *LogRecord)   // WARNING级别日志
cLog.Error(logRecord *LogRecord)     // ERROR级别日志
cLog.Critical(logRecord *LogRecord)  // CRITICAL级别日志
cLog.Fixed(logRecord *LogRecord)     // FIXED级别日志
```

## 日志参数说明  

### 初始化函数参数说明

参阅：[初始化函数参数列表](#初始化函数参数列表)  

### 日志输出函数参数及日志字段列表

日志输出函数传入参数为一个结构体`LogRecord`。

```go
// 允许设置的log内容
type LogRecord struct {
	Message    string    `json:"message,omitempty"`  
	Tag        string    `json:"tag"`
	TraceId    string    `json:"trace_id,omitempty"`
	ExcInfo    string    `json:"exc_info,omitempty,string"`
	Extra      *ExtField `json:"extra,omitempty"`
}
```


| 参数名称          | 日志字段   | 参数说明 | 参数类型 | 必传 | 其它说明 |
| ----------------- | ---------- | -------- | -------- | ---- | ---- |
|         无          | level_name | 日志级别（日志类型分类） | int<br> (10:DEBUG<br> 20:INFO <br>30:WARNING<br> 40:ERROR <br>50:CRITICAL<br>100:FIXED) | 是 | 调用不同级别的日志输出函数时，自动生成。 |
|          无         | log_time   | 日志时间 | string | 是 | 自动生成 |
|            无       | filename   | 文件名 | string | 是 | 自动生成 |
|            无       | module     | 模块名 | string | 是 | 自动生成 |
|              无     | line_no   | 行号 | int | 是 | 自动生成 |
|              无     | func_name  | 函数名 | string | 是 | 自动生成 |
|        logRecord.Message           | message    | 日志内容 | string | 是 | 需手动填写 |
| logRecord.Tag | tag        | 内部日志标签（同一服务允许多个不同的标签，用于区分不同主题） | string | 否 | 需手动填写，默认值为root |
| logRecord.TraceId | trace_id   | 追踪标识（用于追踪服务调用链） | string | 否 | 需手动填写 |
| logRecord.ExcInfo | exc_info   | 异常信息 | string | 否 | 需指定异常或错误信息的对象 |
| 无 | stack_info | 调用堆栈 | string | 否 | exc_info有异常时自动生成 |
| LogRecord.Extra | extra | 扩展字段 | *ExtField | 否 | 可添加任意字段。其中，`type ExtField map[string]interface{}` |
|无|@fluentd_tag|全局日志标签（同一服务使用唯一标签，如聊天服务为dq2_chat）|string|是|由初始化函数InitLogger的loggerName参数指定|
