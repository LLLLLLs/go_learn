//@time:2020/01/20
//@desc:

package mockdata

import (
	"fmt"
	"time"
)

const (
	defaultOnceCount  = 50000
	defaultSleeping   = time.Second * 5
	defaultInsertOnce = 1000
)

type config struct {
	total      int64         // 构建总量
	onceCount  int           // 每次执行插入数量
	insertOnce int           // 每次插入数量
	sleeping   time.Duration // 每次执行后休眠时间
	beginId    int64         // id开始值
}

type Option func(c *config)

func OnceCountOpt(onceCount int) Option {
	return func(c *config) {
		c.onceCount = onceCount
	}
}

func SleepingOpt(sleeping time.Duration) Option {
	return func(c *config) {
		c.sleeping = sleeping
	}
}

func BeginIdOpt(beginId int64) Option {
	return func(c *config) {
		c.beginId = beginId
	}
}

func InsertOnceOpt(insertOnce int) Option {
	return func(c *config) {
		c.insertOnce = insertOnce
	}
}

func NewConf(total int64, options ...Option) *config {
	conf := &config{
		total:      total,
		onceCount:  defaultOnceCount,
		sleeping:   defaultSleeping,
		insertOnce: defaultInsertOnce,
	}
	for _, opt := range options {
		opt(conf)
	}
	if conf.onceCount%conf.insertOnce != 0 {
		panic(fmt.Sprintf("onceCount必须被insertOnce整除,onceCount=%d,insertOnce=%d", conf.onceCount, conf.insertOnce))
	}
	return conf
}
