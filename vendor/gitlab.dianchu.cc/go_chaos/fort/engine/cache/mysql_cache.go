package cache

import (
	"time"

	"gitlab.dianchu.cc/go_chaos/fort/utils"
)

type MySQLCacher struct {
	operation *CacheAsidePattern
}

// Redis 地址 端口 密码 db,缓存模式
func NewMySQLCacher(host, port, password string, db int) *MySQLCacher {
	redisClient := utils.NewRedisManager(host, port, password, db)
	cacher := &MySQLCacher{operation: &CacheAsidePattern{expiration: 600 * time.Second, rc: redisClient}}
	return cacher
}




func (c *MySQLCacher) GetCache(database, table string, pkVals []interface{}, model interface{}) (bool, int64, error) {
	return c.operation.GetCache(database, table, pkVals, model)
}

func (c *MySQLCacher) UpdateCache(database, table string, pkVals []interface{}, data []byte) error {
	return c.operation.UpdateCache(database, table, pkVals, data)
}

func (c *MySQLCacher) CacherCheck() bool {
	return c.operation != &CacheAsidePattern{}
}

func (c *MySQLCacher) UpdateCacheVersion(database, table string, pkVal interface{}, lock bool) error {
	return c.operation.UpdateCacheVersion(database, table, pkVal, lock)
}
