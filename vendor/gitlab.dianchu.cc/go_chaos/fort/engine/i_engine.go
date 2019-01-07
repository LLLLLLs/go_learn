package engine

import (
	"context"
	"database/sql"

	"gitlab.dianchu.cc/go_chaos/fort/engine/cache"
)

type IDALEngine interface {
	SetMySQLCacher(cacher *cache.MySQLCacher)
	MySQLCacher() *cache.MySQLCacher
	GetDBName() string
	SetEncoding(encodingType string, triggerSize int)
	DALDisTransaction(ctx context.Context, data []DALTransaction) error
	DALTransaction(ctx context.Context, data []CmdData) error
	DALQuery(ctx context.Context, data CmdData) (int64, []map[string]interface{}, error)
}

type ISQLEngine interface{
	DataSourceName() string
	Drivers() string
	DB() *sql.DB
	SetMySQLCacher(cacher *cache.MySQLCacher)
	MySQLCacher() *cache.MySQLCacher
}
