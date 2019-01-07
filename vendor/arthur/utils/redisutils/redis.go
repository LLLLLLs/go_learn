package redisutils

import (
	"arthur/conf"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

const (
	defaultDB = 0
)

var (
	client *redis.Client
)

func Init(c conf.RedisConf) {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	rwTimeout := time.Duration(c.RWTimeout) * time.Second

	client = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     c.Password,
		DB:           defaultDB,
		PoolSize:     500,
		ReadTimeout:  rwTimeout,
		WriteTimeout: rwTimeout,
		DialTimeout:  time.Duration(c.ConnTimeout) * time.Second,
	})
}

// 获取redis连接,index 代表redis中的数据库索引,获取后用defer RedisClient.CloseConn()归还连接
func GetClient() *redis.Client {
	if client == nil {
		panic("must Init() before GetClient")
	}
	return client
}

func Close() {
	if client != nil {
		client.Close()
	}
}
