/*
Author : Haoyuan Liu
Time   : 2018/5/28
*/
package infra

import (
	"arthur/app/base"
	"arthur/sdk/uuid"
	"github.com/asaskevich/EventBus"
	"gitlab.dianchu.cc/goutil/dcapi.v2"
)

type Context interface {
	dcapi.Context
	TraceId() string
	//事件总线
	EventBus() EventBus.Bus
}

// 游戏会话，实例的生命周期为一个请求
type context struct {
	dcapi.Context
	traceId string
	bus     EventBus.Bus
}

func NewContext(manager *dcapi.Manager, traceId string) Context {
	return &context{
		Context: dcapi.NewContext(manager),
		traceId: traceId,
	}
}

func (ctx *context) EventBus() EventBus.Bus {
	if ctx.bus == nil {
		ctx.bus = EventBus.New()
	}
	return ctx.bus
}

func (ctx *context) TraceId() string {
	return ctx.traceId
}

func MockContext() Context {
	traceId := uuid.GetUuidBySystem(1)[0]
	ctx := NewContext(dcapi.New(), traceId)
	ctx.Store().SetImmutable(base.TRACE_ID, traceId)
	ctx.Store().SetImmutable(base.IP, "")
	return ctx
}
