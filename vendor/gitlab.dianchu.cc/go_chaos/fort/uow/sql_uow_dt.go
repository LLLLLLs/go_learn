package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	"gitlab.dianchu.cc/go_chaos/fort/engine"
	"gitlab.dianchu.cc/go_chaos/fort/event"
	"gitlab.dianchu.cc/go_chaos/fort/factory"
	"gitlab.dianchu.cc/go_chaos/fort/tools/syslog"
)

// 分布式事务提交
func DTCommit(ctx context.Context, uows ...UoW) error {
	var (
		uowAgentType   reflect.Type
		transaction    = make(map[string][]*factory.SQLInfo)
		err            error
		commandsHandle commandsHandleFunc
	)
	disCommitInit := func(ag event.EventAgent) error {
		if uowAgentType == nil {
			uowAgentType = reflect.TypeOf(ag)
			switch uowAgentType {
			case reflect.TypeOf(&event.DALAgent{}):
				commandsHandle = dalCommandsHandle
			case reflect.TypeOf(&event.SQLAgent{}):
				commandsHandle = sqlCommandsHandle
			default:
				return errors.New("type error:" + fmt.Sprint(uowAgentType))
			}
		} else if uowAgentType != reflect.TypeOf(ag) {
			return errors.New("There are different engine types ")
		}
		return nil
	}

	var ag event.EventAgent
	for _, uow := range uows {
		ag = uow.GetAgent()
		if err = disCommitInit(ag); err != nil {
			return err
		}
		for _, sqlInfo := range ag.GetCommands() {
			transaction[ag.GetSourceName()] = append(transaction[ag.GetSourceName()], sqlInfo)
		}
	}
	return commandsHandle(ctx, transaction, ag)
}

type commandsHandleFunc func(context.Context, map[string][]*factory.SQLInfo, event.EventAgent) error

func dalCommandsHandle(ctx context.Context, transaction map[string][]*factory.SQLInfo, ag event.EventAgent) error {
	var (
		dbSourceName string
		sqlInfo      []*factory.SQLInfo
	)
	cmdData := make(map[string][]engine.CmdData)
	dalDB := ag.GetEngine().(engine.IDALEngine)
	for dbSourceName, sqlInfo = range transaction {
		for _, cmd := range sqlInfo {
			data := engine.CmdData{
				Statement: cmd.Sql,
				Args:      cmd.Args,
			}
			cmdData[dbSourceName] = append(cmdData[dbSourceName], data)
		}
	}
	if len(cmdData) == 1 {
		syslog.FortLog.ShowLog(syslog.DEBUG, cmdData[dbSourceName])
		return dalDB.DALTransaction(ctx, cmdData[dbSourceName])
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
		return dalDB.DALDisTransaction(ctx, sqlData)
	}
}

func sqlCommandsHandle(ctx context.Context, transaction map[string][]*factory.SQLInfo, ag event.EventAgent) error {
	for dbSourceName, sqlInfo := range transaction {
		var (
			tx  *sql.Tx
			err error
		)
		sqlEngine, err := engine.NewSQLEngine(ag.GetEngine().(engine.ISQLEngine).Drivers(), dbSourceName)
		if tx, err = sqlEngine.DB().BeginTx(ctx, nil); err != nil {
			return err
		}
		for _, cmd := range sqlInfo {
			_, err = tx.ExecContext(ctx, cmd.Sql, cmd.Args...)
			if err != nil {
				errBack := tx.Rollback()
				if errBack != nil {
					return errBack
				}
				return err
			}
		}
		if err = tx.Commit(); err != nil {
			errBack := tx.Rollback()
			if errBack != sql.ErrTxDone && errBack != nil {
				return errBack
			}
			return err
		}
	}
	return nil
}
