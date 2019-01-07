package zookeeper

import (
	"arthur/utils/panicutils"
	"github.com/samuel/go-zookeeper/zk"
	"gitlab.dianchu.cc/DevOpsGroup/goutils/exp"
	"sync"
	"time"
)

//Zookeeper节点对象
type Node struct {
	path string
	sync.Mutex
	stat *zk.Stat
}

// 新增节点
func NewNode(path string) INode {
	return &Node{
		path,
		sync.Mutex{},
		nil,
	}
}

func (n *Node) Path() string {
	return n.path
}

func (n *Node) Stat() *zk.Stat {
	return n.stat
}

func (n *Node) SetStat(stat *zk.Stat) {
	n.stat = stat
}

func (n *Node) Get() (data []byte, err error) {
	var stat *zk.Stat
	data, stat, err = Client.Get(n.path)
	if err != nil {
		return nil, err
	}
	n.SetStat(stat)
	return data, err
}

func (n *Node) Set(data []byte) error {
	if n.Stat() == nil {
		_, getErr := n.Get()
		if getErr != nil {
			return getErr
		}
	}
	stat, err := Client.Set(n.Path(), data, n.Stat().Version)
	if err != nil {
		return err
	}
	n.SetStat(stat)
	return nil
}

//被监控的节点对象
type NodeWatched struct {
	INode
	ech      <-chan zk.Event // 事件通道
	callback WatchCallback   // 节点变动后执行的回调
	quit     chan QuitSign   // 退出信号
}

func newWatchedNode(n INode, cb WatchCallback) INodeWatched {
	return &NodeWatched{
		n,
		make(<-chan zk.Event),
		cb,
		make(chan QuitSign),
	}
}

func (w *NodeWatched) GetNode() INode {
	return w.INode
}

func (w *NodeWatched) Ech() <-chan zk.Event {
	return w.ech
}

func (w *NodeWatched) Quit() chan QuitSign {
	return w.quit
}

func (w *NodeWatched) SendQuit(sign QuitSign) {
	w.quit <- sign
}

func (w *NodeWatched) DoCallback(b []byte) {
	if w == nil && w.callback == nil {
		panic("未开始回调或回调函数为nil")
	}
	w.callback(b)
}

// 设置回调函数
func (w *NodeWatched) SetCallback(cb WatchCallback) {
	w.callback = cb
}

func (w *NodeWatched) GetW() (data []byte, err error) {
	var stat *zk.Stat
	data, stat, w.ech, err = Client.GetW(w.Path())
	w.SetStat(stat)
	if err != nil {
		return nil, err
	}
	return data, err
}

func watchNode(w INodeWatched) {
	var (
		data []byte
		err  error
	)
	path := w.Path()
	exp.Try(func() {
		logger.Info("开始监控节点：", path)
		for {
			select {
			case event := <-w.Ech():
				if Client.State() != zk.StateHasSession {
					if Client.State() != zk.StateDisconnected {
						Client.Close()
					}
					login()
					data, err = w.GetW()
					panicutils.OkOrPanic(err)
					if err != nil {
						logger.Error(err)
						continue
					}
					w.DoCallback(data)
				}

				// 若值变换，则更新数据
				if event.Type == zk.EventNodeDataChanged {
					data, err = w.GetW()
					panicutils.OkOrPanic(err)
					if err != nil {
						logger.Error(err)
						continue
					}
					logger.Info("执行zk回调：", path, "节点值：", string(data))
					w.DoCallback(data)
					logger.Info("结束zk回调")
				}
			case <-w.Quit():
				logger.Info("节点监控已停止：", path)
				return
			}
		}
	}, func(ex exp.Exception) {
		logger.Error(ex.Message)
		time.Sleep(3 * time.Second)
	})
}
