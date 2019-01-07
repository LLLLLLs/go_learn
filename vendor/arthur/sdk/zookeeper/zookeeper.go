/*
zookeeper包用于建立全局单例的zk连接
使用前必须执行 InitInfo() 初始化连接

可以使用 StartWatch() 与 StopWatch() 将INode 与 IWatchedNode互相转换
*/
package zookeeper

import (
	"arthur/utils/panicutils"
	"github.com/samuel/go-zookeeper/zk"
	"sync"
	"time"
	"arthur/utils/log"
)

const timeout = time.Second * 10

var (
	Client *Conn // zk连接单例

	logger = log.Logger
)

type (
	// 节点值变动时执行的回调
	WatchCallback func([]byte)

	// 退出监控的信号
	QuitSign struct{}

	// ZK 连接
	Conn struct {
		zk.Conn
		auth string // 登录信息
	}

	// 节点
	INode interface {
		sync.Locker
		Path() string // 节点路径

		Get() ([]byte, error) // 获取节点值
		Set([]byte) error     // 设置节点值

		Stat() *zk.Stat   // 获取节点状态
		SetStat(*zk.Stat) // 设置节点状态
	}

	// 被监控的节点
	INodeWatched interface {
		INode
		GetNode() INode // 获得该Watcher的INode

		GetW() ([]byte, error) // 获取节点值并触发回调、持续监控

		Ech() <-chan zk.Event // zk事件

		Quit() chan QuitSign // 退出监控的信号
		SendQuit(QuitSign)   // 发送退出信号

		DoCallback([]byte)         // 值变更时执行的回调
		SetCallback(WatchCallback) // 修改回调
	}
)

// 初始化zk配置
func Init(zkHosts []string, zkAuth string) *Conn {
	conn, _, err := zk.Connect(zkHosts, timeout)
	panicutils.OkOrPanic(err)
	Client = &Conn{*conn, zkAuth}
	login()
	return Client
}

func CreateNode(path string, b []byte) error{
	_, err := Client.Create(path, b, 0, zk.WorldACL(zk.PermAll))
	return err
}

// 登录zk
func login() {
	if Client == nil {
		panic("登录前必须之前Init()")
	}
	Client.AddAuth("digest", []byte(Client.auth))
}

// 开始监控节点值
func StartWatch(n INode, cb WatchCallback) (w INodeWatched, initData []byte ){
	w = newWatchedNode(n, cb)
	var err error
	initData, err = w.GetW()
	panicutils.OkOrPanic(err)
	go watchNode(w)
	return w, initData
}

// 停止监控节点
func StopWatch(w INodeWatched) INode {
	w.SendQuit(QuitSign{})
	n := w.GetNode()
	return n
}
