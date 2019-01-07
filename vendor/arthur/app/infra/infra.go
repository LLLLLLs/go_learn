/*
Author : Haoyuan Liu
Time   : 2018/6/21
*/
package infra

import (
	"arthur/conf"
	"arthur/env"
	"arthur/sdk/behaviorloguow"
	"arthur/sdk/chat"
	"arthur/sdk/dclog"
	"arthur/sdk/mail"
	"arthur/sdk/push"
	"arthur/sdk/session"
	"arthur/sdk/uuid"
	"arthur/sdk/zookeeper"
	"arthur/utils/redisutils"

	"arthur/notify"
	"gitlab.dianchu.cc/go_chaos/fort/tools/syslog"
)

// 初始化应用
func Init() {
	InitSdk()
	//初始化数据连接
	InitConn()
	//初始化redis
	redisutils.Init(conf.Config.Redis)
	//初始化通知事件
	notify.Init(env.ZK_ROOT)
}

// 初始化第三方服务
func InitSdk() {
	// 设置zookeeper关闭
	zookeeper.Init(env.ZK_HOST, env.ZK_AUTH)

	// 初始化设置
	conf.Init(env.ZK_ROOT)
	cfg := conf.Config

	if behaviorloguow.Toggle {
		behaviorloguow.Init( // 初始化行为日志
			cfg.BehaviorLog.ZKConn.Servers,
			cfg.BehaviorLog.ZKConn.FlumePath,
			cfg.BehaviorLog.ZKConn.ConfPath,
			conf.IsMode(conf.DEBUG),
		)
	}

	// 初始化Sdk
	dclog.Init(cfg.DCLog.Host, cfg.DCLog.Port, env.GAME_NAME)
	session.Init(
		conf.Config.Session.Timeout,
		conf.Config.Session.Method,
		conf.Config.Session.HttpAddr,
	)
	chat.Init(
		cfg.Chat.Timeout,
		cfg.Chat.PageLimit,
		cfg.Chat.Method,
		cfg.Chat.GrpcAddr,
		cfg.Chat.HttpAddr,
		cfg.Chat.Version,
		cfg.Chat.IsPush,
	)
	push.Init(
		cfg.Push.Timeout,
		cfg.Push.Method,
		cfg.Push.HttpAddr,
		cfg.Push.GrpcAddr,
	)
	mail.InitMail(
		cfg.Mail.Timeout,
		cfg.Mail.Method,
		cfg.Mail.GrpcAddr,
		cfg.Mail.HttpAddr,
		cfg.Mail.Version,
		cfg.Mail.IsPush,
	)
	uuid.Init(cfg.Uuid.HttpAddr, cfg.Uuid.PoolSize)
	syslog.FortLog.SetShowLog(false)
}
