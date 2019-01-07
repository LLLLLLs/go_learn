/*
* 配置管理集成，负责实时拉取zk上的配置并对其进行解析，按照定义的格式返回
 */
package flumesdk

import (
	"encoding/json"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"os"
	"strings"
	"sync"
	"time"
)

// 配置管理
type ConfigManager struct {
	zkClient  *zk.Conn    //zk连接
	servers   []string    //zkhost
	flumePath string      //flume路径
	confPath  string      //flume配置路径
	Host      *FlumeHosts //flume节点
	Conf      *Config     //sdk配置
	Data      []byte
}

// 对zk上的配置进行监听，如果配置发生变更，更新配置
func (C *ConfigManager) watch(event <-chan zk.Event) {
	var (
		e         zk.Event
		err       error
		nextEvent <-chan zk.Event
	)

	select {
	case e = <-event:
		switch e.Type {
		case zk.EventNodeDataChanged:
			C.Data, _, nextEvent, err = C.zkClient.GetW(e.Path)
			if err != nil {
				//todo 添加日志组件
				fmt.Printf("%v", err)
			}
			data := string(C.Data)
			fmt.Println(data)

			if strings.Contains(data, "flume_hosts") {
				var host FlumeHosts
				json.Unmarshal(C.Data, &host)
				C.Host = &host
			} else {
				var conf Config
				json.Unmarshal(C.Data, &conf)
				C.Conf = &conf
			}
			go C.watch(nextEvent)
		case zk.EventSession:
			if e.State == zk.StateDisconnected || e.State == zk.StateExpired {
				C.initClient()
				//todo 添加日志组件
				fmt.Printf("EventSession stare:%s,path:%s", e.State.String(), e.Path)
			}
		}
	case <-time.After(time.Second * 5):
		if C.zkClient.State() == zk.StateDisconnected {
			C.initClient()
			fmt.Println("StateDisconnected!")
			return
		}
		go C.watch(event)
	}
}

// 关闭zk
func (C *ConfigManager) Close() {
	C.zkClient.Close()
}

// 从环境变量拉取zk参数
func (C *ConfigManager) InitEnvVars() {
	var (
		err           error
		envExists     bool
		sysEnvServers string
	)
	// ZK_SERVERS
	sysEnvServers, envExists = os.LookupEnv("ZK_SERVERS")
	if !envExists {
		//todo 添加日志组件
		fmt.Printf("zk_server is not set in env")
		panic(err)
	}
	C.servers = strings.Split(sysEnvServers, ",")
	// FLUME_HOST_PATH
	C.flumePath, envExists = os.LookupEnv("FLUME_HOST_PATH")
	if !envExists {
		fmt.Printf("flume_host_path is not set")
		panic(err)
	}
	// FLUME_CONF_PATH
	C.confPath, envExists = os.LookupEnv("FLUME_CONF_PATH")
	if !envExists {
		fmt.Printf("flume_conf_path is not set")
		panic(err)
	}
}

// 初始化zk的连接，拉取初始配置并开启监听
func (C *ConfigManager) initClient() {
	var (
		flumeEvent <-chan zk.Event
		confEvent  <-chan zk.Event
		err        error
	)

	C.zkClient, _, err = zk.Connect(C.servers, time.Second)
	if err != nil {
		panic(err)
	}

	C.Data, _, flumeEvent, err = C.zkClient.GetW(C.flumePath)
	if err != nil {
		panic(err)
	}
	var host FlumeHosts
	json.Unmarshal(C.Data, &host)
	C.Host = &host
	go C.watch(flumeEvent)

	C.Data, _, confEvent, err = C.zkClient.GetW(C.confPath)
	if err != nil {
		panic(err)
	}
	var conf Config
	json.Unmarshal(C.Data, &conf)
	C.Conf = &conf
	go C.watch(confEvent)
}

//获取flumeHosts配置，获取flume节点
func (C *ConfigManager) GetHosts() *FlumeHosts {
	return C.Host
}

// 获取文件缓存路径
func (C *ConfigManager) GetFlumePath() string {
	return C.Conf.FlumeEvent
}

//获取单发配置
func (C *ConfigManager) GetSingle() SingleSend {
	return C.Conf.SingleSend
}

// 获取批发配置
func (C *ConfigManager) GetBatch() BatchSend {
	return C.Conf.BatchSend
}

// 获取扫描配置
func (C *ConfigManager) GetScan() ScanConf {
	return C.Conf.ScanConf
}

// 获取批发日志
func (C *ConfigManager) GetBatchTable() string {
	return strings.Join(C.Conf.BatchSendTableName, ":")
}

var myConfigManager *ConfigManager
var once sync.Once

// 创建并初始化conf并返回,配置管理为单例模式
func newConfigManager() *ConfigManager {
	if myConfigManager != nil {
		return myConfigManager
	}
	once.Do(func() {
		myConfigManager = new(ConfigManager)
		myConfigManager.InitEnvVars()
		myConfigManager.initClient()
	})
	return myConfigManager
}

func newConfigManagerParam(zkServer []string, flumePath, confPath string) *ConfigManager {
	if myConfigManager != nil {
		return myConfigManager
	}
	myConfigManager = new(ConfigManager)
	myConfigManager.servers = zkServer
	myConfigManager.flumePath = flumePath
	myConfigManager.confPath = confPath
	myConfigManager.initClient()
	return myConfigManager
}
