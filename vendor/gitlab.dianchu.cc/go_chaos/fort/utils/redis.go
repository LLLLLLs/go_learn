package utils

import (
	"bytes"
	"time"

	"github.com/go-redis/redis"
)

type RedisManager struct {
	client   *redis.Client // 内部已经实现了连接池与自我熔断机制
	host     string
	port     string
	db       int
	password string
}

func NewRedisManager(host, port, password string, db int) *RedisManager {
	var addr bytes.Buffer
	addr.WriteString(host)
	addr.WriteString(":")
	addr.WriteString(port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr.String(),
		Password: password,
		DB:       db,
	})
	return &RedisManager{client: client, host: host, port: port, db: db, password: password}
}

//exp 时间为0 代表用不过期
func (r RedisManager) Set(key string, value interface{}, exp time.Duration) error {
	return r.client.Set(key, value, exp).Err()
}

func (r RedisManager) Get(key string) (bool, []byte, error) {
	var (
		data []byte
		err  error
		res  *redis.StringCmd
	)
	res = r.client.Get(key)
	if len(res.Val()) == 0 {
		return false, nil, nil
	}
	if data, err = res.Bytes(); err != nil {
		return false, nil, nil
	}
	return true, data, err
}

func (r RedisManager) HSet(key, field string, value interface{}) (bool, error) {
	return r.client.HSet(key, field, value).Result()
}

func (r RedisManager) HGet(key, field string) (bool, string) {
	data := r.client.HGet(key, field).Val()
	if len(data) == 0 {
		return false, ""
	}
	return true, data
}
