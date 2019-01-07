package event

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"gitlab.dianchu.cc/go_chaos/fort/engine"
	"gitlab.dianchu.cc/go_chaos/fort/engine/cache"
	"gitlab.dianchu.cc/go_chaos/fort/factory"
	"gitlab.dianchu.cc/go_chaos/fort/tools/syslog"
)

type SQLAgent struct {
	SQLBaseAgent
	engine engine.ISQLEngine
	driver string
	db     *sql.DB
	tx     *sql.Tx
	cacher *cache.MySQLCacher
}

//NewSQLAgent(ctx, "mysql", "username:password@tcp(host:port)/database?charset=utf8mb4", false)
//isDTS :True 支持分布式事务,False 则是一个单独的事务
func NewSQLAgent(ctx context.Context, sqlEngine engine.ISQLEngine, isDTS bool) (*SQLAgent, error) {
	var (
		agent *SQLAgent
		err   error
	)

	defer func() {
		if err != nil {
			syslog.FortLog.ShowLog(syslog.ERROR, err.Error())
		}
	}()

	agent = new(SQLAgent)
	agent.engine = sqlEngine
	agent.db = agent.engine.DB()
	agent.driver = agent.engine.Drivers()
	agent.cacher = agent.engine.MySQLCacher()
	if isDTS {
		if err = agent.initCommandMemory(ctx, sqlEngine.DataSourceName()); err != nil {
			return nil, err
		}
	} else {
		agent.sourceName = sqlEngine.DataSourceName()
		if agent.tx, err = agent.db.BeginTx(ctx, nil); err != nil {
			return nil, err
		}
	}
	return agent, err
}

func (agent *SQLAgent) Submit(ctx context.Context, sqlQuery *factory.SQLInfo) error {
	if agent.tx == nil {
		syslog.FortLog.ShowLog(syslog.DEBUG, sqlQuery.Sql, sqlQuery.Args)
		return agent.submit(sqlQuery)
	} else {
		//todo 此处也需要指出缓存怎么办
		syslog.FortLog.ShowLog(syslog.DEBUG, sqlQuery.Sql, sqlQuery.Args)
		_, err := agent.tx.ExecContext(ctx, sqlQuery.Sql, sqlQuery.Args...)
		if err != nil {
			fmt.Println(":::::::",err)
			errBack := agent.tx.Rollback()
			if errBack != nil {
				return errBack
			}
			return err
		}
	}
	return nil
}

func (agent *SQLAgent) Extend(commands *[]*factory.SQLInfo) error {
	return agent.extend(commands)
}

func (agent *SQLAgent) commit() error {
	var (
		err     error
		errBack error
	)
	err = agent.tx.Commit()
	if err != nil {
		errBack = agent.tx.Rollback()
		if errBack != nil {
			return errBack
		}
		return err
	}
	return nil
}

func (agent *SQLAgent) Commit(ctx context.Context) error {
	defer func() {
		agent.transaction = &sync.Map{}
	}()

	if agent.tx != nil {
		return agent.commit()
	}

	var (
		tx        *sql.Tx
		cmd       *factory.SQLInfo
		errBack   error
		commitErr error
	)

	agent.transaction.Range(func(sourceName, sqlAtom interface{}) bool {
		sqlEngine, err := engine.NewSQLEngine(agent.driver, sourceName.(string))
		if err != nil {
			commitErr = err
			return false
		}
		if tx, err = sqlEngine.DB().BeginTx(ctx, nil); err != nil {
			commitErr = err
			return false
		}
		for _, cmd = range sqlAtom.(*factory.SQLAtom).CmdList {
			_, err = tx.ExecContext(ctx, cmd.Sql, cmd.Args...)
			if err != nil {
				errBack = tx.Rollback()
				if errBack != nil {
					commitErr = errBack
					return false
				}
				commitErr = err
				return false
			}
		}
		if err = tx.Commit(); err != nil {
			errBack = tx.Rollback()
			if errBack != sql.ErrTxDone && errBack != nil {
				commitErr = errBack
				return false
			}
			commitErr = err
			return false
		}
		return true
	})
	return commitErr
}

func (agent *SQLAgent) GetSourceName() string {
	return agent.sourceName
}

func (agent *SQLAgent) SetSourceName(sourceName string) {
	agent.sourceName = sourceName
}

func (agent *SQLAgent) CanCache() bool {
	if agent.cacher == nil {
		return false
	}
	return agent.cacher.CacherCheck()
}

func (agent *SQLAgent) GetEngine() interface{} {
	return agent.engine
}
