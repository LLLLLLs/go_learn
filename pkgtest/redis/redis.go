// @author: lls
// @date: 2021/8/31
// @desc:

package redis

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	client *redis.Client
	once   = &sync.Once{}
)

func Client() *redis.Client {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr: ":6379",
			DB:   1,
		})
	})
	return client
}
