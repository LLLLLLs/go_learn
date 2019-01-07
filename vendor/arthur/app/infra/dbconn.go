/*
Author : Haoyuan Liu
Time   : 2018/6/6
*/
package infra

import (
	"arthur/app/base"
	"arthur/conf"
	"arthur/utils/panicutils"
	"gitlab.dianchu.cc/go_chaos/fort/engine"
	"arthur/app/model/center"
	"arthur/env"
	"fmt"
	"sync"
)

type connMethod int

const (
	dalConn    connMethod = 1 //通过DAL服务
	directConn connMethod = 2 //直接连接SQL数据库

	DriverName = "mysql"
)

var (
	//连接数据库的方式
	dbConn = dalConn

	sqlConnMap = sqlConnWithLock{
		mux:     &sync.RWMutex{},
		connMap: make(map[base.ASID]engine.ISQLEngine),
	}
	sqlCenterConn engine.ISQLEngine
	dalConnMap    = dalConnWithLock{
		mux:     &sync.RWMutex{},
		connMap: make(map[base.ASID]engine.IDALEngine),
	}
	dalCenterConn engine.IDALEngine

	dbRefMap = make(map[conf.DSN]*dbRef)
)

type sqlConnWithLock struct {
	mux     *sync.RWMutex
	connMap map[base.ASID]engine.ISQLEngine
}

type dalConnWithLock struct {
	mux     *sync.RWMutex
	connMap map[base.ASID]engine.IDALEngine
}

type dbRef struct {
	db    engine.ISQLEngine
	times int
}

func getDBConn() connMethod {
	if conf.IsMode(conf.TEST) {
		return directConn
	}
	return dbConn
}

func InitConn() {
	switch getDBConn() {
	case dalConn:
		initDALConn()
	case directConn:
		initSQLConn()
	default:
		panic("wrong connection method")
	}
}

func AddDBConn(as base.ASID, dbName string) {
	host := conf.Config.Database.DalZkHost
	auth := conf.Config.Database.DalZkAuth
	path := conf.Config.Database.DalZkPath
	switch getDBConn() {
	case dalConn:
		dalConnMap.mux.Lock()
		defer dalConnMap.mux.Unlock()
		client, err := engine.NewDALEngine(dbName, path, host, auth)
		panicutils.OkOrPanic(err)
		dalConnMap.connMap[as] = client
	case directConn:
		sqlConnMap.mux.Lock()
		defer sqlConnMap.mux.Unlock()
		sqlConnMap.connMap[as] = initSQL(conf.GetUri(dbName))
	default:
		panic("wrong connection method")
	}
}

func Close() {
	sqlConnMap.mux.Lock()
	defer sqlConnMap.mux.Unlock()
	for _, d := range sqlConnMap.connMap {
		d.DB().Close()
	}
}

func initSQLConn() {
	sqlCenterConn = initSQL(conf.CenterUri())
	dbConfig := getDbConf()
	sqlConnMap.mux.Lock()
	defer sqlConnMap.mux.Unlock()
	for k, v := range dbConfig {
		sqlConnMap.connMap[base.ASID(k)] = initSQL(conf.DSN(v))
	}
}

func initDALConn() {
	var err error
	host := conf.Config.Database.DalZkHost
	auth := conf.Config.Database.DalZkAuth
	path := conf.Config.Database.DalZkPath
	// 初始化中心数据库
	dalCenterConn, err = engine.NewDALEngine(conf.Config.Database.CenterDB, path, host, auth)
	panicutils.OkOrPanic(err)
	dbConfig := getDbConf()
	// 初始化区服数据库
	dalConnMap.mux.Lock()
	defer dalConnMap.mux.Unlock()
	for asid, dbName := range dbConfig {
		client, err := engine.NewDALEngine(dbName, path, host, auth)
		panicutils.OkOrPanic(err)
		dalConnMap.connMap[base.ASID(asid)] = client
	}
}

func getDbConf() map[base.ASID]string {
	var servers = make([]center.ServerInfo, 0)
	err := QueryCenter().All(MockContext(), &servers)
	panicutils.OkOrPanic(err)
	dbConfig := make(map[base.ASID]string)
	for _, server := range servers {
		as := fmt.Sprintf("%d%s%d", server.AppId, env.REDIS_KEY_SEP, server.ServerId)
		switch getDBConn() {
		case dalConn:
			dbConfig[base.ASID(as)] = server.DbName
		case directConn:
			dbConfig[base.ASID(as)] = string(conf.GetUri(server.DbName))
		default:
			panic("wrong connection method")
		}
	}
	return dbConfig
}

func initSQL(d conf.DSN) engine.ISQLEngine {
	var (
		db  engine.ISQLEngine
		ref *dbRef
		err error
		ok  bool
	)
	if ref, ok = dbRefMap[d]; ok {
		ref.times += 1
		return ref.db
	}

	db, err = engine.NewSQLEngine(DriverName, string(d))
	if err != nil {
		panic(err)
	}
	ref = &dbRef{
		db,
		1,
	}
	dbRefMap[d] = ref
	return db
}
