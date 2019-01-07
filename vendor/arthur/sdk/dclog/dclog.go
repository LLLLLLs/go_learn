/*
package: actions
author: Wenhao Shan
date: 2018/05/08
desc: chaos log
*/
package dclog

import (
	"arthur/utils/log"
	"github.com/sirupsen/logrus"
	glo "gitlab.dianchu.cc/chaos_go_sdk/log_out_sdk_go"
	"strconv"
	"strings"
)

var IsRemote = false

func Init(host string, port int, loggerName string) {
	logIp := host
	logPort := strconv.Itoa(port)
	glo.Logger.InitLogger(false, IsRemote, strings.ToUpper(log.Logger.Level().String()), logIp, logPort,
		"", loggerName)
}

func doLog(level log.Level, message, tag, traceId string, excMap glo.ExtField, excInfo string, extend ...interface{}) {
	entry := log.Logger.WithFields(logrus.Fields(excMap))
	switch level {
	case log.DebugLevel:
		entry.Debug(message)
	case log.InfoLevel:
		entry.Info(message)
	case log.WarnLevel:
		entry.Warning(message)
	case log.ErrorLevel:
		entry.Error(message)
	case log.FatalLevel:
		entry.Fatal(message)
	}
	record := logRecord(message, tag, traceId, excMap, excInfo)
	glo.Logger.Log(levelToGlo[level], record, extend...)
}

/*
	Debug 调试级别的日志, 针对细粒度信息事件
		message: 	日志内容
		tag: 		日志标签，用于区分输入日志的模块
		traceId:	追踪标识
		excMap:		拓展字段
*/
func Debug(message, tag, traceId string, excMap glo.ExtField) {
	doLog(log.DebugLevel, message, tag, traceId, excMap, "")
}

/*
	Info 程序正常运行级别的日志, 针对粗粒度级别
		message: 	日志内容
		tag: 		日志标签，用于区分输入日志的模块
		traceId:	追踪标识
		excMap:		拓展字段
*/
func Info(message, tag, traceId string, excMap glo.ExtField) {
	doLog(log.InfoLevel, message, tag, traceId, excMap, "")
}

/*
	Warning 错误级别的日志, 表明发生错误事件, 但不影响系统的继续运行
		message: 	日志内容
		tag: 		日志标签，用于区分输入日志的模块
		traceId:	追踪标识
		excMap:		拓展字段
*/
func Warning(message, tag, traceId string, excMap glo.ExtField) {
	doLog(log.WarnLevel, message, tag, traceId, excMap, "")
}

/*
	SetError
		message: 	日志内容
		tag: 		日志标签，用于区分输入日志的模块
		traceId:	追踪标识
		excMap:		拓展字段
		excInfo:	栈信息
*/
func Error(message, tag, traceId string, excMap glo.ExtField, excInfo string) {
	doLog(log.ErrorLevel, message, tag, traceId, excMap, excInfo)
}

/*
	Fatal 崩溃级别的日志, 表明发生了引起服务器停止运行级别的错误
		message: 	日志内容
		tag: 		日志标签，用于区分输入日志的模块
		traceId:	追踪标识
		excMap:		拓展字段
		excInfo:	栈信息
*/
func Fatal(message, tag, traceId string, excMap glo.ExtField, excInfo string) {
	doLog(log.FatalLevel, message, tag, traceId, excMap, excInfo)
}

/*
	执行日志记录
		message: 	日志内容
		tag: 		日志标签
		traceId:	追踪标识
		excMap:		拓展字段
		excInfo:	栈信息
*/
func logRecord(message, tag, traceId string, excMap glo.ExtField, excInfo string) *glo.LogRecord {
	if excMap == nil {
		excMap = make(glo.ExtField)
	}
	logRec := &glo.LogRecord{
		Message: message,
		Tag:     tag,
		TraceId: traceId,
		ExcInfo: excInfo,
		Extra:   &excMap,
	}
	return logRec
}
