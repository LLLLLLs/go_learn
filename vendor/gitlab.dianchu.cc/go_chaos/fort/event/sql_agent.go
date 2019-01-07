package event

import (
	"context"
	"errors"
	"math"
	"reflect"
	"sync"

	"gitlab.dianchu.cc/go_chaos/fort/engine"
	"gitlab.dianchu.cc/go_chaos/fort/factory"
	"gitlab.dianchu.cc/go_chaos/fort/tools/syslog"
	"gitlab.dianchu.cc/go_chaos/fort/utils"
)

type EventAgent interface {
	Commit(ctx context.Context) error
	Submit(ctx context.Context, sqlQuery *factory.SQLInfo) error
	Extend(commands *[]*factory.SQLInfo) error
	GetCommands() []*factory.SQLInfo
	SetSourceName(sourceName string)
	GetSourceName() string
	GetEngine() interface{}
	CanCache() bool
}

type SQLBaseAgent struct {
	sourceName  string
	transaction *sync.Map
	cmdlistMux  *sync.Mutex
}

func (agent *SQLBaseAgent) initCommandMemory(ctx context.Context, sourceName string) error {
	var (
		ok bool
	)
	agent.cmdlistMux = &sync.Mutex{}
	agent.transaction = new(sync.Map)
	atom := new(factory.SQLAtom)
	agent.sourceName = sourceName
	if _, ok = ctx.Value(utils.TRACE_ID).(string); !ok {
		errStr := "The transaction has not trace id! "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return errors.New(errStr)
	}
	atom.CmdList = make([]*factory.SQLInfo, 0, utils.COMMAND_CAP)
	agent.transaction.Store(sourceName, atom)
	return nil
}

func (agent *SQLBaseAgent) remalloc(atom *factory.SQLAtom, size int) {
	var (
		number int
		extra  []*factory.SQLInfo
	)
	number = len(atom.CmdList)
	extra = make([]*factory.SQLInfo, number, int(math.Ceil(float64(number+size)/utils.COMMAND_CAP))*utils.COMMAND_CAP)
	for i := 0; i < number; i++ {
		extra[i] = atom.CmdList[i]
	}
	atom.CmdList = extra
}

func (agent *SQLBaseAgent) submit(sqlQuery *factory.SQLInfo) error {
	defer agent.cmdlistMux.Unlock()
	agent.cmdlistMux.Lock()
	atom := new(factory.SQLAtom)
	v, _ := agent.transaction.LoadOrStore(agent.sourceName, new(factory.SQLAtom))
	atom = v.(*factory.SQLAtom)

	cmdCount := len(atom.CmdList) + 1
	if cmdCount > cap(atom.CmdList) {
		agent.remalloc(atom, 1)
	}
	atom.CmdList = atom.CmdList[:cmdCount]
	atom.CmdList[cmdCount-1] = sqlQuery
	return nil
}

func (agent *SQLBaseAgent) extend(commands *[]*factory.SQLInfo) error {
	defer agent.cmdlistMux.Unlock()
	agent.cmdlistMux.Lock()
	var (
		atom *factory.SQLAtom
	)
	if v, ok := agent.transaction.Load(agent.sourceName); ok {
		atom = v.(*factory.SQLAtom)
	}
	cmdCount := len(atom.CmdList)
	extCount := len(*commands)
	total := cmdCount + extCount
	if total > cap(atom.CmdList) {
		agent.remalloc(atom, extCount)
	}
	atom.CmdList = atom.CmdList[:total]
	for i := cmdCount; i < total; i++ {
		atom.CmdList[i] = (*commands)[i-cmdCount]
	}
	return nil
}

func (agent *SQLBaseAgent) GetCommands() []*factory.SQLInfo {
	if v, ok := agent.transaction.Load(agent.sourceName); ok {
		return v.(*factory.SQLAtom).CmdList
	}
	return nil
}

func (agent *SQLBaseAgent) SetSourceName(sourceName string) {
	agent.sourceName = sourceName
}

func NewAgent(ctx context.Context, dbEngine interface{}, args ...interface{}) (EventAgent, error) {
	t := reflect.TypeOf(dbEngine)
	if t.Kind() != reflect.Ptr {
		return nil, errors.New(t.Kind().String() + "is not ptr type! ")
	}
	switch reflect.TypeOf(dbEngine) {
	case reflect.TypeOf(&engine.DALEngine{}):
		return NewDALAgent(ctx, dbEngine.(*engine.DALEngine))
	case reflect.TypeOf(&engine.SQLEngine{}):
		isDTS := false
		if len(args) > 1 {
			if i, ok := args[0].(bool); ok {
				isDTS = i
			}
		}
		return NewSQLAgent(ctx, dbEngine.(*engine.SQLEngine), isDTS)
	default:
		return nil, errors.New("error engine")
	}
}
