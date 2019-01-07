package event

import (
	"context"
	"reflect"
	"sync"

	"gitlab.dianchu.cc/go_chaos/fort/engine"
	"gitlab.dianchu.cc/go_chaos/fort/engine/cache"
	"gitlab.dianchu.cc/go_chaos/fort/factory"
	"gitlab.dianchu.cc/go_chaos/fort/tools/syslog"
)

type DALAgent struct {
	SQLBaseAgent
	engine engine.IDALEngine
	cacher *cache.MySQLCacher
}

func NewDALAgent(ctx context.Context, dalEngine engine.IDALEngine) (*DALAgent, error) {
	var (
		dal *DALAgent
		err error
	)
	dal = new(DALAgent)
	if err = dal.initCommandMemory(ctx, dalEngine.GetDBName()); err != nil {
		return nil, err
	}
	dal.engine = dalEngine
	dal.cacher = dalEngine.MySQLCacher()
	return dal, nil
}

func (dal *DALAgent) Submit(ctx context.Context, sqlQuery *factory.SQLInfo) error {
	return dal.submit(sqlQuery)
}

func (dal *DALAgent) Extend(commands *[]*factory.SQLInfo) error {
	return dal.extend(commands)
}

func (dal *DALAgent) Commit(ctx context.Context) error {
	defer func() {
		dal.transaction = &sync.Map{}
	}()
	var (
		err     error
		dbName  string
		sqlAtom *factory.SQLAtom
	)
	cmdData := make(map[string][]engine.CmdData)
	// 收集此次提交涉及的行级数据
	idCache := make(map[string]map[string][]interface{})

	dal.transaction.Range(func(name, atom interface{}) bool {
		dbName = name.(string)
		sqlAtom = atom.(*factory.SQLAtom)
		idCache[dbName] = make(map[string][]interface{})
		for _, cmd := range sqlAtom.CmdList {
			data := engine.CmdData{
				Statement: cmd.Sql,
				Args:      cmd.Args,
			}
			cmdData[dbName] = append(cmdData[dbName], data)
			//单条语句涉及多个主键的情况(Add方法) CACHE
			if pkValue := reflect.ValueOf(cmd.PrimaryKeyValue); pkValue.Kind() == reflect.Slice {
				//todo 此处反射需要校验，
				for i := 0; i < pkValue.Len(); i++ {
					idCache[dbName][cmd.TableName] = append(idCache[dbName][cmd.TableName], pkValue.Index(i))
				}
			} else {
				idCache[dbName][cmd.TableName] = append(idCache[dbName][cmd.TableName], cmd.PrimaryKeyValue)
			}
		}
		return true
	})

	updateCache := func(lock bool) {
		for database, val0 := range idCache {
			for table, val1 := range val0 {
				for _, pkVal := range val1 {
					dal.cacher.UpdateCacheVersion(database, table, pkVal, lock)
				}
			}
		}
	}

	if dal.CanCache() {
		// 锁定此次提交涉及的行级数据,期间直接读取数据数据。
		updateCache(true)
	}

	if len(cmdData) == 1 {
		syslog.FortLog.ShowLog(syslog.DEBUG, cmdData[dbName])
		err = dal.engine.DALTransaction(ctx, cmdData[dbName])
	} else {
		var sqlData []engine.DALTransaction
		for dbName, cmdInfo := range cmdData {
			data := engine.DALTransaction{
				DB:   dbName,
				Data: cmdInfo,
			}
			sqlData = append(sqlData, data)
		}
		syslog.FortLog.ShowLog(syslog.DEBUG, sqlData)
		err = dal.engine.DALDisTransaction(ctx, sqlData)
	}
	if dal.CanCache() {
		if err == nil {
			updateCache(false)
		}
	}
	return err
}

func (dal *DALAgent) GetSourceName() string {
	return dal.sourceName
}

func (dal *DALAgent) SetSourceName(sourceName string) {
	//dal.db.SetDBName(sourceName)
	dal.sourceName = sourceName
}

func (dal *DALAgent) CanCache() bool {
	if dal.cacher == nil {
		return false
	}
	return dal.cacher.CacherCheck()
}

func (dal *DALAgent) GetEngine() interface{} {
	return dal.engine
}
