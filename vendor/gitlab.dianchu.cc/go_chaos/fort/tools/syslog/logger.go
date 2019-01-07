package syslog

import (
	"fmt"
	"runtime"

	cLog "gitlab.dianchu.cc/chaos_go_sdk/log_out_sdk_go"
)

const (
	ERROR   = cLog.ERROR
	INFO    = cLog.INFO
	WARNING = cLog.WARNING
	DEBUG   = cLog.DEBUG
)

type ILogger interface {
	Debug(logRecord *cLog.LogRecord)
	Info(logRecord *cLog.LogRecord)
	Warning(logRecord *cLog.LogRecord)
	Error(logRecord *cLog.LogRecord)
	SetLevel(level int)
}

var FortLog = &SysLog{
	logger: cLog.Logger,
}

type SysLog struct {
	log     func(logRecord *cLog.LogRecord)
	logger  ILogger
	level   int
	showLog bool
}

// 输出日志
func (l *SysLog) ShowLog(level int, msg ...interface{}) {
	if l.level >= level && l.showLog {
		_, file, line, _ := runtime.Caller(2)
		if l.log == nil {
			l.SetLevel(level)
		}
		l.log(&cLog.LogRecord{
			Message: fmt.Sprintf("%v", msg),
			Extra: &cLog.ExtField{
				"File": file,
				"Line": line,
			},
		})
	}
}

// 设置日志等级
func (l *SysLog) SetLevel(level int) {
	l.level = level
	switch l.level {
	case DEBUG:
		l.logger.SetLevel(DEBUG)
		l.log = l.logger.Debug
	case INFO:
		l.logger.SetLevel(INFO)
		l.log = l.logger.Info
	case WARNING:
		l.logger.SetLevel(WARNING)
		l.log = l.logger.Warning
	case ERROR:
		l.logger.SetLevel(ERROR)
		l.log = l.logger.Error
	}
}

// 可接受外部系统日志对象
func (l *SysLog) SetLogger(logger ILogger) {
	l.logger = logger
}

// 是否开启日志输出
func (l *SysLog) SetShowLog(b bool) {
	l.showLog = b
}
