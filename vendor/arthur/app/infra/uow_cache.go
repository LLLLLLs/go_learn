package infra

import (
	"arthur/utils/redisutils"
	goctx "context"
	"github.com/go-redis/redis"
)

type UoWCache interface {
	redis.Cmdable
	Commit(goctx.Context) error
}

type uowCache struct {
	redis.Pipeliner
}

func newUowCache() *uowCache {
	pipe := redisutils.GetClient().Pipeline()
	return &uowCache{
		Pipeliner: pipe,
	}
}

func (u uowCache) Commit(ctx goctx.Context) error {
	_, err := u.Exec()
	return err
}
