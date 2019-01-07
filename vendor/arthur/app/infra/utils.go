package infra

import (
	"arthur/app/base"
	goctx "context"
	"arthur/utils/errors"
)

func GetTraceId(ctx goctx.Context) string {
	traceId, ok := ctx.Value(base.TRACE_ID).(string)
	if !ok {
		panic("must init TRACE_ID in context")
	}
	return traceId
}

func GetIP(ctx goctx.Context) string {
	ip, ok := ctx.Value(base.IP).(string)
	if !ok {
		panic("must init IP in context")
	}
	return ip
}

func GetASID(ctx goctx.Context) (asid base.ASID, err error) {
	var ok bool

	value := ctx.Value(base.APPSERVER_ID)
	if value == nil {
		return "", errors.New("cannot find appserver in ctx")
	}
	asid, ok = value.(base.ASID)
	if !ok {
		return "", errors.New("cannot find appserver in ctx")
	}
	return asid, nil
}

func NewCtxWithTraceID(traceID string) goctx.Context{
	return goctx.WithValue(goctx.Background(), base.TRACE_ID, traceID)
}