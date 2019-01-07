package engine

import (
	"database/sql"
	"sync"

	"gitlab.dianchu.cc/go_chaos/fort/engine/cache"
)

//Apache Ignite/GridGain: https://github.com/amsokol/ignite-go-client
//Apache Avatica/Phoenix: https://github.com/apache/calcite-avatica-go
//ClickHouse (uses native TCP interface): https://github.com/kshvakov/clickhouse
//ClickHouse (uses HTTP API): https://github.com/mailru/go-clickhouse
//Couchbase N1QL: https://github.com/couchbase/go_n1ql
//DB2 LUW and DB2/Z with DB2-Connect: https://bitbucket.org/phiggins/db2cli (Last updated 2015-08)
//DB2 LUW (uses cgo): https://github.com/asifjalil/cli
//Firebird SQL: https://github.com/nakagami/firebirdsql
//MS ADODB: https://github.com/mattn/go-adodb
//MS SQL Server (pure go): https://github.com/denisenkom/go-mssqldb
//MS SQL Server (uses cgo): https://github.com/minus5/gofreetds
//MySQL: https://github.com/ziutek/mymysql [*]
//MySQL: https://github.com/go-sql-driver/mysql/ [*]
//ODBC: https://bitbucket.org/miquella/mgodbc (Last updated 2016-02)
//ODBC: https://github.com/alexbrainman/odbc
//Oracle: https://github.com/mattn/go-oci8
//Oracle: https://gopkg.in/rana/ora.v4
//Oracle: https://gopkg.in/goracle.v2
//QL: http://godoc.org/github.com/cznic/ql/driver
//Postgres (pure Go): https://github.com/lib/pq [*]
//Postgres (uses cgo): https://github.com/jbarham/gopgsqldriver
//Postgres (pure Go): https://github.com/jackc/pgx [**]
//SAP HANA (pure go): https://github.com/SAP/go-hdb
//Snowflake (pure Go): https://github.com/snowflakedb/gosnowflake
//SQLite (uses cgo): https://github.com/mattn/go-sqlite3 [*]
//SQLite (uses cgo): https://github.com/gwenn/gosqlite - Supports SQLite dynamic data typing
//SQLite (uses cgo): https://github.com/mxk/go-sqlite
//SQLite: (uses cgo): https://github.com/rsc/sqlite
//Sybase SQL Anywhere: https://github.com/a-palchikov/sqlago
//Vitess: https://godoc.org/github.com/youtube/vitess/go/vt/vitessdriver
//YQL (Yahoo! Query Language): https://github.com/mattn/go-yql

var (
	directStorer = &directEngineStore{
		Store: make(map[string]*SQLEngine),
	}
)

//并发安全,保存**sql.Open之后的db对象**,防止重复创建相同地址的db对象
type directEngineStore struct {
	sync.RWMutex
	Store map[string]*SQLEngine
}

type SQLEngine struct {
	dataSourceName string
	drivers        string
	db             *sql.DB //已经自己实现连接池了
	mysqlCacher    *cache.MySQLCacher
}

func (store *directEngineStore) getEngine(sourceName string) *SQLEngine {
	store.RLock()
	sqlEngine, ok := store.Store[sourceName]
	store.RUnlock()
	if ok {
		return sqlEngine
	}
	return nil
}

func (store *directEngineStore) addEngine(sourceName string, sqlEngine *SQLEngine) {
	store.Lock()
	store.Store[sourceName] = sqlEngine
	store.Unlock()
}

func NewSQLEngine(driverName, dataSourceName string) (ISQLEngine, error) {
	var (
		s *SQLEngine
	)
	if s = directStorer.getEngine(driverName + dataSourceName); s != nil {
		if s.db.Ping() == nil {
			return s, nil
		}
	}
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	s = &SQLEngine{
		dataSourceName: dataSourceName,
		drivers:        driverName,
		db:             db,
	}
	directStorer.addEngine(driverName+dataSourceName, s)
	return s, err
}

func (dal *SQLEngine) DataSourceName() string {
	return dal.dataSourceName
}

func (dal *SQLEngine) Drivers() string {
	return dal.drivers
}

func (dal *SQLEngine) DB() *sql.DB {
	return dal.db
}

func (dal *SQLEngine) SetMySQLCacher(cacher *cache.MySQLCacher) {
	dal.mysqlCacher = cacher
}

func (dal *SQLEngine) MySQLCacher() *cache.MySQLCacher {
	return dal.mysqlCacher
}
