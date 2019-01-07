/*
Author : Haoyuan Liu
Doc	   : 封装logrus库
*/
package log

import "github.com/sirupsen/logrus"

type Fields map[string]interface{}

type Level uint32

func (level Level) String() string {
	switch level {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case panicLevel:
		return "panic"
	}

	return "unknown"
}

const (
	panicLevel = Level(logrus.PanicLevel)
	FatalLevel = Level(logrus.FatalLevel)
	ErrorLevel = Level(logrus.ErrorLevel)
	WarnLevel  = Level(logrus.WarnLevel)
	InfoLevel  = Level(logrus.InfoLevel)
	DebugLevel = Level(logrus.DebugLevel)
)
