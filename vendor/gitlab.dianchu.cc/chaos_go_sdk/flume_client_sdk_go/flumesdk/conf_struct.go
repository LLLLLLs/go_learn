/*
* 定义zk配置结构,用于解析zk配置
 */

package flumesdk

import (
	"fmt"
)

type FlumeNode struct {
	Node  string `json:"node"`
	Host  string `json:"host"`
	Port  int    `json:"port"`
	State int    `json:"state"`
}

//获取flumeHosts配置，获取flume节点
func (F *FlumeNode) GetHostPort() string {
	var host string
	host = fmt.Sprintf("%s:%d", F.Host, F.Port)
	return host
}

type FlumeHosts struct {
	Hosts []FlumeNode `json:"flume_hosts"`
}

func (FH *FlumeHosts) GetValidNode() []string {
	var hosts []string
	for _, node := range FH.Hosts {
		if node.State == 1 {
			hosts = append(hosts, node.GetHostPort())
		}
	}
	return hosts
}

type AppId struct {
	Dq1  []int `json:"dq_1"`
	Dq2  []int `json:"dq_2"`
	Tang []int `json:"tang"`
	Test []int `json:"test"`
}

type FilePath struct {
	FlumeEvent string `json:"flume_event"`
}
type SendConf struct {
	SingleSend `json:"single_send"`
	BatchSend  `json:"batch_send"`
}
type SingleSend struct {
	Timeout int
}
type BatchSend struct {
	Timeout    int
	BatchCount int `json:"batch_count"`
}

type Control struct {
	MaxWait  int `json:"max_wait"`
	ConnMin  int `json:"conn_min"`
	PoolSize int `json:"pool_size"`
}

func (C *Control) getMaxWait() int {
	if C.MaxWait == 0 {
		C.MaxWait = 200
	}
	return C.MaxWait
}

func (C *Control) getConnMin() int {
	if C.ConnMin == 0 {
		C.ConnMin = 20
	}
	return C.ConnMin
}

func (C *Control) getPoolSize() int {
	if C.PoolSize == 0 {
		C.PoolSize = 50
	}
	return C.PoolSize
}

type ScanConf struct {
	Interval   int `json:"interval"`
	FileCount  int `json:"file_count"`
	QueueCount int `json:"queue_count"`
	ScanQueue  int `json:"scan_queue"`
	ScanFile   int `json:"scan_file"`
}

type Config struct {
	AppId              `json:"app_id"`
	FilePath           `json:"file_path"`
	SendConf           `json:"send_conf"`
	ScanConf           `json:"scan_conf"`
	BatchSendTableName []string `json:"batch_send_table_name"`
	Control            `json:"control"`
}
