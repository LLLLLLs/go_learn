package uow

import (
	"context"

	"gitlab.dianchu.cc/chaos_go_sdk/flume_client_sdk_go/flumesdk"
	"gitlab.dianchu.cc/go_chaos/fort/event"
	"gitlab.dianchu.cc/go_chaos/fort/factory"
)

type UoW interface {
	Add(ctx context.Context, model interface{}) error
	Save(ctx context.Context, model interface{}) error
	Modify(ctx context.Context, model interface{}, fields map[string]interface{}) error
	Remove(ctx context.Context, model interface{}) error
	Submit(ctx context.Context, sqlInfo *factory.SQLInfo) error
	Commit(ctx context.Context) error
	GetAgent() event.EventAgent
	Log(eventLog flumesdk.EventLog) error
	SetSourceName(sourceName string)
}
