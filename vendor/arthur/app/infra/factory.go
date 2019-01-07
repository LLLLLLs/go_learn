/*
Author : Haoyuan Liu
Time   : 2018/6/4
*/
package infra

import (
	"arthur/utils/panicutils"
	fortfty "gitlab.dianchu.cc/go_chaos/fort/factory"
)

var factory *fortfty.SQLFactory

type Schema fortfty.Schema

func RegisterFactory(modelSet ...interface{}) {
	err := factory.Register(modelSet...)
	panicutils.OkOrPanic(err)
}

func init(){
	factory = fortfty.NewSQLFactory()
}