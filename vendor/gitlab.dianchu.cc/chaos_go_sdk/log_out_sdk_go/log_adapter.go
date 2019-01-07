package log_output

import (
	"fmt"
	"os"

	"encoding/json"
	"errors"
	"io"
	"strconv"
	"sync"
	"time"
	// json "github.com/json-iterator/go"
)

//==================syslog==========================
var syslogWriter SysLogHandle

//==================================================
/*
FIXED 等级用于记录固定的信息，比如程序启动，关闭等等。
CRITICAL 级别以上会输出栈信息
*/
const (
	DEBUG    = 10
	INFO     = 20
	WARNING  = 30
	ERROR    = 40
	CRITICAL = 50
	FIXED    = 100

	openBrace   = '{'
	closeBrace  = '}'
	comma       = ','
	doubleQuote = '"'
	singleQuote = '\''
	colon       = ':'
	newLine     = '\n'
)

const RootLoggerName = "root_logger"
const DefaultTag = "root"

var GlobleStdLock = &sync.Mutex{}

var NameToLevel = map[string]int{
	"DEBUG":    DEBUG,
	"INFO":     INFO,
	"WARNING":  WARNING,
	"ERROR":    ERROR,
	"CRITICAL": CRITICAL,
	"FIXED":    FIXED,
}

var LevelToName = map[int]string{
	DEBUG:    "DEBUG",
	INFO:     "INFO",
	WARNING:  "WARNING",
	ERROR:    "ERROR",
	CRITICAL: "CRITICAL",
	FIXED:    "FIXED",
}
var NoMatchLogLevel = errors.New("can't match log level")

// log 主体
type CustomLogger struct {
	Name         string
	Level        int
	FixedFlag    bool // 是否启用 FIXED 日志输出,默认是true
	mu           *sync.Mutex
	out          io.Writer
	Tag          []byte
	CloserWriter *SysLogHandle
	FluentdTag   string
}

// 日志输出的字段，true表示可以在拓展字段中覆盖他
var recordField = map[string]bool{
	"level":          false,
	"log_time":       true,
	"filename":       false,
	"moudle":         false,
	"line_no":        false,
	"func_name":      false,
	"message":        false,
	"tag":            false,
	"trace_id":       false,
	"exc_info":       false,
	"stack_info":     false,
	"generated_time": false,
}

// 3 allocs/op
type ExtField map[string]interface{}

// 允许设置的log内容
type LogRecord struct {
	Message string    `json:"message,omitempty"`
	Tag     string    `json:"tag"`
	TraceId string    `json:"trace_id,omitempty"`
	ExcInfo string    `json:"exc_info,omitempty,string"`
	Extra   *ExtField `json:"extra,omitempty"`
}

// 是否直接调用log的标志,用来设置stack_skip层数的,自定义以便识别.
type LogCallFlag int

// LogLevel 设置日志等级
func LogLevel(level int) error {
	_, ok := LevelToName[level]
	if !ok {
		return NoMatchLogLevel
	}
	if Logger.Level > 0 {
		Logger.Level = level
	}
	return nil
}

// GetTime 获取时间（毫秒级）
func GetTime() string {
	now := time.Now()
	// %Y-%m-%d %H:%M:%S
	// 2016 -> Y 年
	// 01   -> m 月
	// 02   -> d 日
	// 15   -> H 小时
	// 04   -> M 分钟
	// 05   -> S 秒
	// .000 ->   毫秒
	return now.Format("2006-01-02 15:04:05.000")
}

// GetTime 获取时间（毫秒级）
func GetTimeRFC3339() string {
	now := time.Now()
	// %Y-%m-%d %H:%M:%S
	// 2016 -> Y 年
	// 01   -> m 月
	// 02   -> d 日
	// 15   -> H 小时
	// 04   -> M 分钟
	// 05   -> S 秒
	// .000 ->   毫秒
	return now.Format("2006-01-02T15:04:05Z07:00")
}

// IsEnableLog 是否允许打印日志
func (logger *CustomLogger) isEnableLog(level int) bool {
	//logRecordNew := setFuncInfo(&logRecord,2)
	return (level >= logger.Level) && (logger.FixedFlag || level < FIXED)
	//return level >= logger.Level
}

// SetWriter 设置日志记录容器，默认是os.stdout
func (logger *CustomLogger) SetWriter(writer []io.Writer) {
	// 允许配置多个writer
	if writer != nil {
		logger.out = io.MultiWriter(writer...)
	}
}

// InitLogger 设置日志输出到标志输出
/*
系统日志地址变更理解：
	再次manager.go中Logger对象的InitLogger()方法会重新初始化一个Logger并且把老的那个Logger放到CloserWriter,并且调用其Close对象关闭上次的连接
*/
func (logger *CustomLogger) InitLogger(toStd bool, toSyslog bool, logLevel string, syslogIP string, syslogPort string, logMode string, loggerName string) (err error) {
	// loggerName is Fluentd_tag
	// 环境变量优先级最高
	envToStdout := os.Getenv("LOG_TO_STDOUT")
	envToSyslog := os.Getenv("LOG_TO_SYSLOG")
	envIP := os.Getenv("LOG_SERVER_IP")
	envPort := os.Getenv("LOG_SERVER_PORT")
	envLevel := os.Getenv("LOG_OUT_LEVEL")

	envLogMode := os.Getenv("LOG_MODE")

	// 检查环境变量
	if envLogMode != "" {
		logMode = envLogMode
	}

	if envToStdout == "YES" {
		toStd = true
	} else if envToStdout == "NO" {
		toStd = false
	}

	if envToSyslog == "YES" {
		toSyslog = true
	} else if envToSyslog == "NO" {
		toSyslog = false
	}

	if envIP != "" {
		syslogIP = envIP
	}

	if envPort != "" {
		syslogPort = envPort
	}

	if envLevel != "" {
		logLevel = envLevel
	}

	if logger.Name == RootLoggerName {
		GlobleConf.StdEnable = toStd
		GlobleConf.SyslogEnable = toSyslog
		GlobleConf.LogOutLevel = logLevel
		GlobleConf.ServerIp = syslogIP
		GlobleConf.ServerPort = syslogPort
		GlobleConf.LoggerName = loggerName
		GlobleConf.LogMode = logMode
	}

	if loggerName == "" && GlobleConf.LoggerName != "" {
		loggerName = GlobleConf.LoggerName
	}
	logger.FluentdTag = loggerName

	writers := []io.Writer{}
	var oldSyslog *SysLogHandle
	if toSyslog {
		my_sysl, err = Dial("tcp", syslogIP+":"+syslogPort, syslogLevM[logLevel])
		if err != nil {
			panic(err)
		}
		//当日志地址更新的时候把旧地址的writer给关闭
		if logger.CloserWriter != nil {
			oldSyslog = logger.CloserWriter
		}
		logger.CloserWriter = my_sysl
		writers = append(writers, my_sysl)
	}
	if toStd {
		// writers = append(writers, GetLockWriter(os.Stdout, GlobleStdLock))
		writers = append(writers, os.Stdout)
	}
	logger.SetLevel(defaultLevM[logLevel])
	logger.SetWriter(writers)
	if oldSyslog != nil {
		oldSyslog.Close()
	}
	return nil
}

// SetLevel 设置日志输出等级
func (logger *CustomLogger) SetLevel(Level int) {
	if _, ok := LevelToName[Level]; ok {
		logger.Level = Level
	} else {
		logger.Level = INFO
	}
}

// SetDefaultTag 设置默认TAG
func (logger *CustomLogger) SetDefaultTag(tag string) {
	if tag != "" {
		logger.Tag = EncodeString(tag, false)
	} else {
		logger.Tag = nil
	}
}

// SetFluentdTag 设置 Fluentd Tag
func (logger *CustomLogger) SetFluentdTag(fluentdTag string) {
	logger.FluentdTag = fluentdTag
}

// SetFixedFlag 是否启用 FIXED 日志输出,默认是true
func (logger *CustomLogger) SetFixedFlag(flag bool) {
	logger.FixedFlag = flag
}

// WriterClose  关闭Writer
func (logger *CustomLogger) WriterClose() {
	if logger.CloserWriter != nil {
		logger.CloserWriter.Close()
	}
}

// Log 日志记录，手动写入bytes,效率更快，有待完整测试
func (logger *CustomLogger) Log(level int, logRecord *LogRecord, extend ...interface{}) {
	if !logger.isEnableLog(level) {
		return
	}
	// 输出数据缓冲区
	// lgr := GetlogRecoderHandle()

	data := GetBytesBuffer()
	// defer PutBytesBuffer(data)

	// // 1 allocs/op
	stackSkip := LogCallFlag(2)

	data.WriteByte('{')
	// 设置fluentd_tag
	data.WriteByte('"')
	data.WriteString("@fluentd_tag")
	data.WriteString(`":"`)

	data.WriteString(logger.FluentdTag)
	data.WriteByte('"')

	// 设置log级别 level_name
	data.WriteByte(',')
	data.WriteByte('"')
	data.WriteString("level_name")
	data.WriteString(`":"`)
	data.WriteString(LevelToName[level])
	data.WriteByte('"')

	// 日志生成时间 generated_time
	data.WriteByte(',')
	data.WriteByte('"')
	data.WriteString("generated_time")
	data.WriteString(`":`)
	generatedTime := []byte(fmt.Sprintf("\"%s\"", GetTimeRFC3339()))
	data.Write(generatedTime)

	// 判断是否是直接调用log，非直接调用log的，需要设置一下skip参数，用于栈信息的获取
	for _, v := range extend {
		switch v.(type) {
		case LogCallFlag:
			stackSkip = v.(LogCallFlag)
		}
	}

	// using map
	//设置函数调用信息
	var filename, module, func_name, stack_info string
	var line_no int

	// 设置错误栈信息,level 为 FIXED 时，也不记录
	// 300000	      4680 ns/op	    1200 B/op	       9 allocs/op
	if level >= CRITICAL && level != FIXED {
		// 3600 ns/op 10 allocs/op
		// 1000000	      2583 ns/op	     208 B/op	       1 allocs/op
		stack_info, filename, module, func_name, line_no = CallersWithFirstCallInfo(int(stackSkip) - 1)
	} else {
		filename, module, func_name, line_no = setFuncInfo(int(stackSkip))
	}

	// filename 文件名
	data.WriteByte(',')
	data.WriteByte('"')
	data.WriteString("filename")
	data.WriteString(`":"`)
	data.WriteString(filename)
	data.WriteByte('"')

	// 包名 module
	data.WriteByte(',')
	data.WriteByte('"')
	data.WriteString("module")
	data.WriteString(`":"`)
	data.WriteString(module)
	data.WriteByte('"')

	// 函数名 func_name
	data.WriteByte(',')
	data.WriteByte('"')
	data.WriteString("func_name")
	data.WriteString(`":"`)
	data.WriteString(func_name)
	data.WriteByte('"')

	// 行号 line_no
	data.WriteByte(',')
	data.WriteByte('"')
	data.WriteString("line_no")
	data.WriteString(`":`)
	data.WriteString(strconv.Itoa(line_no))

	// 日志信息 message
	data.WriteByte(',')
	data.WriteByte('"')
	data.WriteString("message")
	data.WriteString(`":`)
	data.Write(EncodeString(logRecord.Message, false))

	// 栈信息 stack_info
	if stack_info != "" {
		data.WriteByte(',')
		data.WriteByte('"')
		data.WriteString("stack_info")
		data.WriteString(`":`)
		// 300000	      4236 ns/op	    1288 B/op	      10 allocs/op
		data.Write(EncodeString(stack_info, false))
	}
	// 错误信息 exc_info
	if logRecord.ExcInfo != "" {
		data.WriteByte(',')
		data.WriteByte('"')
		data.WriteString("exc_info")
		data.WriteString(`":`)
		data.Write(EncodeString(logRecord.ExcInfo, false))
	}

	// 写入trace_id
	if logRecord.TraceId != "" {
		data.WriteByte(',')
		data.WriteByte('"')
		data.WriteString("trace_id")
		data.WriteString(`":"`)
		data.WriteString(logRecord.TraceId)
		data.WriteByte('"')
	}
	// 写入tag
	if logRecord.Tag != "" || logger.Tag != nil {
		data.WriteByte(',')
		data.WriteByte('"')
		data.WriteString("tag")
		data.WriteString(`":`)
		if logRecord.Tag != "" {
			data.Write(EncodeString(logRecord.Tag, false))
		} else if logger.Tag != nil {
			data.Write(logger.Tag)
		}
	}

	// 添加拓展字段的信息
	if logRecord.Extra != nil {
		for k, v := range *logRecord.Extra {
			if _, ok := recordField[k]; ok {
				continue
			}
			data.WriteByte(',')
			switch v.(type) {
			case string:
				data.Write(EncodeString(k, false))
				data.WriteString(`:`)
				data.Write(EncodeString(v.(string), false))
			default:
				data.Write(EncodeString(k, false))
				data.WriteString(`:`)
				tmp, _ := json.Marshal(v)
				data.Write(tmp)
			}
		}
	}

	data.WriteByte('}')
	data.WriteByte('\n')
	// lgr.W = logger.out
	// wp.Serve(lgr)
	if workpoolFlag {
		lgr := GetlogRecoderHandle()
		lgr.W = logger.out
		lgr.Buf = data
		wp.Serve(lgr)
	} else {
		go func() {
			logger.write(data.Bytes())
			PutBytesBuffer(data)
			// PutlogRecoderHandle(lgr)
		}()
	}

}

func (logger *CustomLogger) write(data []byte) {
	// logger.mu.Lock()
	logger.out.Write(data)
	// logger.mu.Unlock()
}

func (logger *CustomLogger) Debug(logRecord *LogRecord) {
	logger.Log(DEBUG, logRecord, LogCallFlag(3))
}

func (logger *CustomLogger) Info(logRecord *LogRecord) {
	logger.Log(INFO, logRecord, LogCallFlag(3))
}

func (logger *CustomLogger) Warning(logRecord *LogRecord) {
	logger.Log(WARNING, logRecord, LogCallFlag(3))
}

func (logger *CustomLogger) Error(logRecord *LogRecord) {
	logger.Log(ERROR, logRecord, LogCallFlag(3))
}
func (logger *CustomLogger) Critical(logRecord *LogRecord) {
	logger.Log(CRITICAL, logRecord, LogCallFlag(3))
}
func (logger *CustomLogger) Fixed(logRecord *LogRecord) {
	logger.Log(FIXED, logRecord, LogCallFlag(3))
}
