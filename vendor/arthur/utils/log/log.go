package log

import (
	"arthur/env"
	"github.com/Robpol86/logrus-custom-formatter"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

var (
	Logger ILogger
)

const (
	//在环境变量中设置改值为要显示的日志tag，tag之间用空格分隔
	//如 ENABLED_TAGS = "value_event sql"则会输出数值事件和uow执行sql的日志信息
	enabledTags = "ENABLED_TAGS"
)

type ILogger interface {
	Level() Level
	SetLevel(level Level)
	SetOutput(output io.Writer)
	WithField(key string, value interface{}) IEntry
	WithFields(fields map[string]interface{}) IEntry
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type IEntry interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type logger struct {
	enabledTags []string
	*logrus.Logger
}

func (l logger) WithField(key string, value interface{}) IEntry {
	if key == "tag" {
		for _, t := range l.enabledTags {
			v := value.(string)
			if v == t || v == "" {
				return l.Logger.WithField(key, value)
			}
		}
	}
	return EmptyEntry{}
}

func (l logger) WithFields(fields map[string]interface{}) IEntry {
	if v, ok := fields["tag"]; ok {
		tagValue := v.(string)
		for _, tag := range l.enabledTags {
			if tagValue == tag {
				return EmptyEntry{}
			}
		}
	}
	return l.Logger.WithFields(fields)
}

func (l logger) SetLevel(level Level) {
	l.Logger.SetLevel(logrus.Level(level))
}

func (l logger) Level() Level {
	return Level(l.Logger.Level)
}

func (l logger) SetOutput(output io.Writer) {
	l.Logger.SetOutput(output)
}

func init() {
	lcf.WindowsEnableNativeANSI(true)
	template := "[%[shortLevelName]s] %-45[message]s%[fields]s\n"
	formatter := lcf.NewFormatter(template, nil)
	l := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: formatter,
		Hooks:     make(logrus.LevelHooks),
	}

	level, err := logrus.ParseLevel(env.LOG_LEVEL)
	if err != nil {
		l.Warning("调试日志等级错误")
		level = logrus.DebugLevel
	}

	l.SetLevel(logrus.Level(level))
	Logger = &logger{
		Logger:      l,
		enabledTags: getenabledTags(),
	}
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Warning(args ...interface{}) {
	Logger.Warning(args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

func Warningf(format string, args ...interface{}) {
	Logger.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}

type EmptyEntry struct{}

func (e EmptyEntry) Debugf(format string, args ...interface{})   {}
func (e EmptyEntry) Infof(format string, args ...interface{})    {}
func (e EmptyEntry) Printf(format string, args ...interface{})   {}
func (e EmptyEntry) Warningf(format string, args ...interface{}) {}
func (e EmptyEntry) Errorf(format string, args ...interface{})   {}
func (e EmptyEntry) Fatalf(format string, args ...interface{})   {}
func (e EmptyEntry) Debug(args ...interface{})                   {}
func (e EmptyEntry) Info(args ...interface{})                    {}
func (e EmptyEntry) Print(args ...interface{})                   {}
func (e EmptyEntry) Warning(args ...interface{})                 {}
func (e EmptyEntry) Error(args ...interface{})                   {}
func (e EmptyEntry) Fatal(args ...interface{})                   {}

func getenabledTags() []string {
	sList := strings.Split(os.Getenv(enabledTags), " ")
	tags := make([]string, len(sList))

	for i, s := range sList {
		tags[i] = strings.TrimSpace(s)
	}
	return tags
}
