/*
Author : Haoyuan Liu
Time   : 2018/6/21
*/
package app

import (
	"arthur/app/info/errors"
	"arthur/app/infra"
	"arthur/app/model"
	"arthur/conf"
	"arthur/sdk/zookeeper"
	"arthur/utils/panicutils"
	"arthur/utils/timeutils"
	"math/rand"
)

//初始化App，App初始化后，该进程内变拥有了一个完备的应用环境
func Init() {
	//初始化系统的随机种子
	rand.Seed(timeutils.Now())
	//加载动态表
	model.LoadDynTable()
	//初始化基础应用
	infra.Init()
	//加载静态配置表
	if !model.LoadStatTable(conf.ProfileDbUri(), &model.ProfCenter) {
		panic(errors.New("load static table failed"))
	}
	//初始化静态键值配置
	model.InitStatKey(model.ProfCenter)
}

//初始化测试App，为测试脚本提供应用环境
func TestInit() {
	//设置系统模式
	conf.SetMode(conf.TEST)
	model.LoadDynTable()
	infra.Init()
	//与Init() 不同，从本地的ProfCenter加载静态配置表（而不是从Mysql里查询），减少了Mysql的压力
	err := model.LoadStatTableInLocal(conf.ProfileDbUri(), &model.ProfCenter)
	panicutils.OkOrPanic(err)
	model.InitStatKey(model.ProfCenter)
}

//关闭App，安全地断开与其他服务的网络连接
func Close() {
	infra.Close()
	zookeeper.Client.Close()
}

func WithTestApp(fn func()) {
	TestInit()
	defer Close()
	fn()
}
