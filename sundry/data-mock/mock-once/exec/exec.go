//@time:2020/01/20
//@desc:

package main

import (
	"flag"
	mockdata "golearn/sundry/data-mock/mock-once"
	"time"
)

var (
	total      = flag.Int64("total", 0, "需要mock的数据总量,必填")
	onceCount  = flag.Int("once", 0, "每次执行插入数据量,默认50000")
	insertOnce = flag.Int("insert", 0, "每次插入条数,默认1000")
	sleeping   = flag.Int("sleep", 0, "每次执行后休眠时间,单位秒,默认5秒")
	begin      = flag.Int64("begin", 0, "数据id起始值,默认根据数据库已有数据计算")
)

func main() {
	flag.Parse()
	opts := make([]mockdata.Option, 0)
	if *onceCount != 0 {
		opts = append(opts, mockdata.OnceCountOpt(*onceCount))
	}
	if *insertOnce != 0 {
		opts = append(opts, mockdata.InsertOnceOpt(*insertOnce))
	}
	if *sleeping != 0 {
		opts = append(opts, mockdata.SleepingOpt(time.Second*time.Duration(*sleeping)))
	}
	if *begin != 0 {
		opts = append(opts, mockdata.BeginIdOpt(*begin))
	}
	conf := mockdata.NewConf(*total, opts...)
	mgr := mockdata.NewMockManager(conf)
	mgr.Mock()
}
