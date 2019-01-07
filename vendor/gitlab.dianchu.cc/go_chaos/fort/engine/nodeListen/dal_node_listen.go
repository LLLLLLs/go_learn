package nodeListen

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"gitlab.dianchu.cc/go_chaos/fort/tools/syslog"
	"gitlab.dianchu.cc/go_chaos/fort/utils"
	"time"
)

type ZKService struct {
	sr     *utils.ServiceRegister
	ech    <-chan zk.Event
	ctx    *context.Context
	ZKHost []string
	ZKAuth []string
	ZKPath string
}

var callback func([]string)

func (srv *ZKService) login() {
	srv.sr = utils.NewServiceRegister(srv.ZKHost, nil, nil)
	ZKAuth := srv.ZKAuth[0] + ":" + fmt.Sprintf("%x", md5.Sum([]byte(srv.ZKAuth[1])))
	srv.sr.ZKClient.AddAuth("digest", []byte(ZKAuth))
}

func (srv *ZKService) getChildren() {
	var (
		data []string
		err  error
	)
	data, _, srv.ech, err = srv.sr.GetChildrenW(srv.ZKPath)
	if err != nil {
		panic(fmt.Sprintf("Zookeeper Error:%v", err))
	}
	callback(data)
}

func (srv *ZKService) NodeListen(ctx *context.Context, fun func([]string)) {
	callback = fun
	srv.login()
	srv.getChildren()
	go srv.eventNodeChildren(ctx)
}

//监听节点的子节点创建
func (srv *ZKService) eventNodeChildren(ctx *context.Context) {
	for {
		select {
		case <-(*ctx).Done():
			return
		case e := <-srv.ech:
			// 子节点新增删除（非数据）
			if e.Type == zk.EventNodeChildrenChanged {
				srv.getChildren()
			}
			// 重连机制
			if e.Type == zk.EventNotWatching && e.State != zk.StateHasSession {
				if e.State == zk.StateUnknown || //未知状态
					e.State == zk.StateExpired || //过期状态
					e.State == zk.StateDisconnected { //断开状态
					time.Sleep(time.Second)
					srv.sr.ZKClient.Close()
					srv.login()
					srv.getChildren()
				}
			}
		}
	}
}

func StartNodeListen(ctx *context.Context, zkHost, zkAuth []string, zkPath string, fun func([]string)) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			syslog.FortLog.ShowLog(syslog.ERROR, fmt.Sprint(panicErr))
			err = errors.New(fmt.Sprint(panicErr))
		}
	}()
	zkSrv := new(ZKService)
	zkSrv.ZKAuth = zkAuth
	zkSrv.ZKHost = zkHost
	zkSrv.ZKPath = zkPath
	zkSrv.NodeListen(ctx, fun)
	return
}
