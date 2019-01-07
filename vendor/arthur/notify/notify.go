/*
Author : Haoyuan Liu
Doc	   : 分布式通知包，可以用该包注册主题，通知所有节点处理事件
*/
package notify

import (
	"arthur/sdk/zookeeper"
	"path"
	"arthur/sdk/dclog"
	"fmt"
	"arthur/utils/panicutils"
	"strconv"
	"time"
)

var (
	nodeName                    = "notify"
	subjectCallback = map[string]func(){}
	watchedNodes map[string]zookeeper.INodeWatched
)

func Init(zkRoot string) {
	notifyRoot := path.Join(zkRoot, nodeName)
	watchedNodes = map[string]zookeeper.INodeWatched{}
	for sub := range subjectCallback {
		nodePath := path.Join(notifyRoot, sub)
		ok, _, err := zookeeper.Client.Exists(nodePath)
		panicutils.OkOrPanic(err)
		if !ok {
			err = zookeeper.CreateNode(nodePath, nowTimeByte())
			panicutils.OkOrPanic(err)
		}
		w, _ := zookeeper.StartWatch(zookeeper.NewNode(nodePath), func(b []byte) {
			subjectCallback[sub]()
		})
		watchedNodes[sub] = w
	}
}

//订阅主题，一旦收到更新提醒，将执行callback
// 一旦执行了Init(), 该函数失效
func RegisterSubject(name string, callback func()) {
	if _, ok := subjectCallback[name]; ok {
		panic(fmt.Sprintf("subject: [%s] has been registered", name))
	}
	subjectCallback[name] = callback
}

//发送某个主题的通知广播
func SendNotify(subject string) {
	checkInit()
	w, ok := watchedNodes[subject]
	if !ok {
		warningMsg := fmt.Sprintf("trying to update subject[%s], but it is not exist", subject)
		dclog.Warning(warningMsg, "sys", "", nil)
	}
	err := w.Set(nowTimeByte())
	if err != nil {
		errMsg := fmt.Sprintf("update subject[%s] failed, err: %s", subject, err.Error())
		dclog.Error(errMsg, "sys", "", nil, err.Error())
	}
}

func checkInit() {
	if watchedNodes == nil {
		panic("must init notify package first")
	}
}

func nowTimeByte() []byte{
	return []byte(strconv.FormatInt(time.Now().UnixNano(), 10))
}
