package log_output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"
)

const configFileSizeLimit = 10 << 20

var (
	Logger     *CustomLogger //控制台消费日志
	GlobleConf globleConf    //全局配置
	my_sysl    *SysLogHandle //syslog.writer
)

type globleConf struct {
	StdEnable    bool   //是否输出到控制台
	SyslogEnable bool   //是否输出到syslog服务器
	LogOutLevel  string //日志输出等级
	ServerIp     string //syslog服务器IP
	ServerPort   string //syslog服务器端口
	LogMode      string //syslog输出模式
	LoggerName   string //logger名称，也即服务标签名，如dq2_chat
}

var syslogLevM map[string]Priority = map[string]Priority{
	"DEBUG":    LOG_DEBUG,
	"INFO":     LOG_INFO,
	"ERROR":    LOG_ERR,
	"WARNING":  LOG_WARNING,
	"CRITICAL": LOG_CRIT,
	"FIXED":    LOG_ALERT,
}

var defaultLevM map[string]int = map[string]int{
	"DEBUG":    DEBUG,
	"INFO":     INFO,
	"WARNING":  WARNING,
	"ERROR":    ERROR,
	"CRITICAL": CRITICAL,
	"FIXED":    FIXED,
}

//============================
// 初始化
func init() {
	// 默认参数
	GlobleConf.StdEnable = true
	GlobleConf.SyslogEnable = false
	GlobleConf.LogOutLevel = "INFO"
	GlobleConf.LoggerName = "log_test"
	Logger = GetLogger(RootLoggerName, DefaultTag)
	Logger.SetWriter([]io.Writer{os.Stdout})
}

var loggerManager = make(map[string]*CustomLogger, 1)

var lock = sync.RWMutex{}

// GetLogger 获取 Logger
func GetLogger(name string, tag string) *CustomLogger {
	lock.RLock()
	logger, ok := loggerManager[name]
	lock.RUnlock()
	if !ok {
		lock.Lock()
		defer lock.Unlock()
		logger = &CustomLogger{
			Level:      defaultLevM[GlobleConf.LogOutLevel],
			FixedFlag:  true,
			mu:         &sync.Mutex{},
			FluentdTag: GlobleConf.LoggerName,
		}
		logger.Name = name
		if logger.Name != RootLoggerName {
			err := logger.InitLogger(
				GlobleConf.StdEnable,
				GlobleConf.SyslogEnable,
				GlobleConf.LogOutLevel,
				GlobleConf.ServerIp,
				GlobleConf.ServerPort,
				GlobleConf.LogMode,
				GlobleConf.LoggerName,
			)
			if err != nil {
				fmt.Println("InitLogger err:", err.Error())
			}
		}

		loggerManager[name] = logger
	}
	if tag != "" {
		logger.SetDefaultTag(tag)
	}
	return logger
}

//json文件读取
func LoadConfig(path string) (config globleConf, err error) {
	config_file, err := os.Open(path)
	if err != nil {
		emit("Failed to open config file '%s': %s\n", path, err)
		return
	}

	fi, _ := config_file.Stat()
	if size := fi.Size(); size > (configFileSizeLimit) {
		emit("config file (%q) size exceeds reasonable limit (%d) - aborting", path, size)
		return // REVU: shouldn't this return an error, then?
	}

	if fi.Size() == 0 {
		emit("config file (%q) is empty, skipping", path)
		return
	}

	buffer := make([]byte, fi.Size())
	_, err = config_file.Read(buffer)
	emit("\n %s\n", buffer)

	buffer, err = StripComments(buffer) //去掉注释
	if err != nil {
		emit("Failed to strip comments from json: %s\n", err)
		return
	}

	buffer = []byte(os.ExpandEnv(string(buffer))) //特殊

	err = json.Unmarshal(buffer, &config) //解析json格式数据
	if err != nil {
		emit("Failed unmarshalling json: %s\n", err)
		return
	}

	return
}

func StripComments(data []byte) ([]byte, error) {
	data = bytes.Replace(data, []byte("\r"), []byte(""), 0) // Windows
	lines := bytes.Split(data, []byte("\n"))                //split to muli lines
	filtered := make([][]byte, 0)

	for _, line := range lines {
		match, err := regexp.Match(`^\s*#`, line)
		if err != nil {
			return nil, err
		}
		if !match {
			filtered = append(filtered, line)
		}
	}

	return bytes.Join(filtered, []byte("\n")), nil
}

func emit(msgfmt string, args ...interface{}) {
	//log.Printf(msgfmt, args...)
}

var workpoolFlag = false

var wp *workerPool

func GetWorkPoolMaxWorks() int {
	return wp.MaxWorkersCount
}

var defaultWorkPoolSize = 200

// ResetWorkPool 设置日志协程池
// 默认设置 200
func ResetWorkPool(maxWorkers int) {
	// 协程池大小一样,就不更新了
	if wp != nil && maxWorkers == wp.MaxWorkersCount {
		return
	}
	newWP := &workerPool{
		WorkerFunc: func(lgr *logRecoderHandle) error {
			_, err := lgr.W.Write(lgr.Buf.Bytes())
			PutlogRecoderHandle(lgr)
			PutBytesBuffer(lgr.Buf)
			return err
		},
		MaxWorkersCount: maxWorkers,
	}
	tmpWP := wp
	wp = newWP
	GlobleStdLock.Lock()
	wp.Start()
	if tmpWP != nil {
		tmpWP.Stop()
	}
	GlobleStdLock.Unlock()
}

// EnableWorkPool 开启workpool
func EnableWorkPool(flag bool, workpoolSize int) {
	workpoolFlag = flag
	if flag {

		if workpoolSize > 0 {
			ResetWorkPool(workpoolSize)
		} else {
			ResetWorkPool(defaultWorkPoolSize)
		}
	} else {
		if wp != nil {
			wp.Stop()
		}
	}
}
