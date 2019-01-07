/*
Author : Haoyuan Liu
Time   : 2018/7/26
*/
package infra

import (
	"arthur/app/base"
	"arthur/app/info/errors"
	"arthur/conf"
	"arthur/env"
	"arthur/sdk/behaviorloguow"
	"arthur/sdk/dclog"
	"arthur/utils/panicutils"
	goctx "context"
	"fmt"
	"gitlab.dianchu.cc/chaos_go_sdk/flume_client_sdk_go/flumesdk"
	"gitlab.dianchu.cc/go_chaos/fort/event"
	fortfactory "gitlab.dianchu.cc/go_chaos/fort/factory"
	fortuow "gitlab.dianchu.cc/go_chaos/fort/uow"
	"reflect"
	"strings"
)

type UoWBase interface {
	//新建数据
	Add(ctx goctx.Context, model interface{}) error
	//保存修改后的model
	Save(ctx goctx.Context, model interface{}) error
	//删除数据
	Remove(ctx goctx.Context, model interface{}) error
	//修改数据
	Modify(ctx goctx.Context, model interface{}, fields map[string]interface{}) error
	//提交所有操作
	Submit(ctx goctx.Context, sqlInfo *fortfactory.SQLInfo) error
	SubmitSQL(ctx goctx.Context, sql string, args []interface{}) error
	//提交所有操作
	Commit(ctx goctx.Context) error
	//获取fort中的uow
	fUoW() fortuow.UoW
	// 增加一条行为日志
	Log(eventLog flumesdk.EventLog) error
	//获取缓存工作单元
	UoWCache() UoWCache
	GetAgent() event.EventAgent
	//Deprecated 使用CommitMany提交多个数据库操作
	SetSourceName(sourceName string)
}

// 区服服务uow
type UoW interface {
	UoWBase
	// 改uow操作的区服
	AppServer() base.AppServer
}

// 中心服务uow
type UoWCenter interface {
	UoWBase
}

type uowBase struct {
	fw    fortuow.UoW //fort的uow
	cache UoWCache    //缓存的uow
}

func newUowBase(fw fortuow.UoW) *uowBase {
	return &uowBase{
		fw: fw,
	}
}

func (w uowBase) Add(ctx goctx.Context, model interface{}) error {
	return w.fw.Add(ctx, model)
}

func (w *uowBase) Save(ctx goctx.Context, model interface{}) error {
	return w.doSave(ctx, model)
}

func (w uowBase) doSave(ctx goctx.Context, model interface{}) error {
	return w.fw.Save(ctx, model)
}

func (w uowBase) Remove(ctx goctx.Context, model interface{}) error {
	return w.fw.Remove(ctx, model)
}

func (w uowBase) Modify(ctx goctx.Context, model interface{}, fields map[string]interface{}) error {
	return w.fw.Modify(ctx, model, fields)
}

func (w uowBase) Submit(ctx goctx.Context, sqlInfo *fortfactory.SQLInfo) error {
	return w.fw.Submit(ctx, sqlInfo)
}

func (w uowBase) SubmitSQL(ctx goctx.Context, sql string, args []interface{}) error {
	return w.fw.Submit(ctx, &fortfactory.SQLInfo{Sql: sql, Args: args})
}

func (w *uowBase) UoWCache() UoWCache {
	if w.cache == nil {
		w.cache = newUowCache()
	}
	return w.cache
}

func (w *uowBase) Commit(ctx goctx.Context) error {
	var err error
	defer w.reset()

	if !conf.IsMode(conf.RELEASE) {
		msg := fmt.Sprintf("uowBase cmd list: %v", GetCMDList(w))
		dclog.Debug(msg, "sql", GetTraceId(ctx), nil)
	}
	err = w.fw.Commit(ctx)
	if err == nil && w.cache != nil {
		e := w.cache.Commit(ctx)
		if e != nil {
			dclog.Error(
				"redis pipe commit error",
				"cache",
				GetTraceId(ctx),
				nil,
				e.Error())
		}
	}
	return err
}

func (w uowBase) GetAgent() event.EventAgent {
	return w.fUoW().GetAgent()
}

func (w uowBase) Log(eventLog flumesdk.EventLog) error {
	return w.fw.Log(eventLog)
}

func (w uowBase) SetSourceName(sourceName string) {
	panic("do not use this method to change db")
}

func (w *uowBase) fUoW() fortuow.UoW {
	return w.fw
}

func (w *uowBase) reset() {
	w.cache = nil
}

type uow struct {
	*uowBase
	as     base.AppServer
	toSave map[string]Model
}

func (w *uow) Save(ctx goctx.Context, model interface{}) error {
	if m, ok := model.(Model); ok {
		name := reflect.TypeOf(m).Elem().Name()
		key := strings.Join([]string{name, m.AppServer().ID().ToString(), m.ID()}, env.REDIS_KEY_SEP)
		w.toSave[key] = m
		return nil
	} else {
		return w.uowBase.Save(ctx, model)
	}
}

func (w *uow) doAllSave(ctx goctx.Context) error {
	for _, m := range w.toSave {
		err := w.doSave(ctx, m)
		if err != nil {
			return err
		}
	}
	w.toSave = map[string]Model{}
	return nil
}

func (w *uow) Commit(ctx goctx.Context) error {
	var err error
	w.doAllSave(ctx)
	err = w.uowBase.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewUow(ctx goctx.Context, as base.AppServer) (UoW, error) {
	var (
		agent event.EventAgent
		err   error
	)
	switch getDBConn() {
	case dalConn:
		dalConnMap.mux.RLock()
		defer dalConnMap.mux.RUnlock()
		client, ok := dalConnMap.connMap[as.ID()]
		if !ok {
			return nil, errors.New("can not find appserverid")
		}
		agent, err = event.NewDALAgent(ctx, client)
		if err != nil {
			return nil, err
		}
	case directConn:
		dalConnMap.mux.RLock()
		defer dalConnMap.mux.RUnlock()
		db, ok := sqlConnMap.connMap[as.ID()]
		if !ok {
			return nil, errors.New("can not find appserverid")
		}
		agent, err = event.NewSQLAgent(ctx, db, true)
		if err != nil {
			return nil, err
		}

	default:
		panic("wrong connection method")
	}

	fw, err := fortuow.NewUoWSQL(ctx, factory, agent, behaviorloguow.EventSender())
	if err != nil {
		return nil, err
	}

	u := newUowBase(fw)
	return &uow{
		uowBase: u,
		as:      as,
		toSave:  map[string]Model{},
	}, nil
}

func (w *uow) AppServer() base.AppServer {
	return w.as
}

func MustNewUow(ctx goctx.Context, as base.AppServer) UoW {
	w, err := NewUow(ctx, as)
	panicutils.OkOrPanic(err)
	return w
}

type uowCenter struct {
	*uowBase
}

func MustNewCenterUow(ctx goctx.Context) UoWCenter {
	w, err := NewCenterUow(ctx)
	panicutils.OkOrPanic(err)
	return w
}

func NewCenterUow(ctx goctx.Context) (UoWCenter, error) {
	var (
		agent event.EventAgent
		err   error
	)
	switch getDBConn() {
	case dalConn:
		agent, err = event.NewDALAgent(ctx, dalCenterConn)
		if err != nil {
			return nil, err
		}
	case directConn:
		agent, err = event.NewSQLAgent(ctx, sqlCenterConn, true)
		if err != nil {
			return nil, err
		}
	default:
		panic("wrong connection method")
	}
	fw, err := fortuow.NewUoWSQL(ctx, factory, agent, behaviorloguow.EventSender())
	if err != nil {
		return nil, err
	}
	w := newUowBase(fw)
	return &uowCenter{w}, err
}

func GetCMDList(uow UoWBase) []string {
	cmds := uow.fUoW().GetAgent().GetCommands()
	cmdList := make([]string, len(cmds))
	for i := range cmds {
		c := cmds[i]
		cmdList[i] = fmt.Sprintf("%s %v", c.Sql, c.Args)
	}
	return cmdList
}

func UowCommitMany(ctx goctx.Context, uows ...fortuow.UoW) error {
	for i := range uows {
		u, ok := uows[i].(*uow)
		if ok {
			err := u.doAllSave(ctx)
			if err != nil {
				return err
			}
		}
	}
	return fortuow.DTCommit(ctx, uows...)
}
