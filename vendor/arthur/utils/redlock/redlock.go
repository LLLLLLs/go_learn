/*
Author : Haoyuan Liu
Time   : 2018/4/19
*/
package redlock

import (
	"arthur/utils/log"
	"arthur/utils/redisutils"
	"github.com/bsm/redis-lock"
	"time"
)

var defaultOptions = lock.Options{
	LockTimeout: 15 * time.Second,
	RetryDelay:  100 * time.Millisecond,
	RetryCount:  0,
}

type Options struct {
	//对key加锁的最大时间
	LockTimeout time.Duration
	//获取锁的重试次数
	RetryCount int
	//每次重试的的间隔时间
	RetryDelay time.Duration
	//加上Prefix后，
	TokenPrefix string
}

type Locker interface {
	//开始加锁，ok为是否拿到锁， err为加锁过程中发生的异常
	Lock() (bool, error)
	Unlock() error
}

type RedLock struct {
	l *lock.Locker
}

//新建锁，Options可为nil
func NewLock(key string, options *Options) Locker {
	client := redisutils.GetClient()
	var o *lock.Options
	if options == nil {
		o = &defaultOptions
	} else {
		o = &lock.Options{
			LockTimeout: options.LockTimeout,
			RetryCount:  options.RetryCount,
			RetryDelay:  options.RetryDelay,
			TokenPrefix: options.TokenPrefix,
		}
	}
	return &RedLock{
		l: lock.New(client, key, o),
	}
}

func (l *RedLock) Lock() (bool, error) {
	ok, err := l.l.Lock()
	if err != nil {
		log.Error("redlock lock failed: ", err)
	}
	return ok, err
}

func (l *RedLock) Unlock() error {
	err := l.l.Unlock()
	if err != nil {
		log.Error("redlock unlock failed: ", err)
	}
	return err
}
