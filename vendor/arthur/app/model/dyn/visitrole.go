/*
Created on 2018-11-06 16:58:42
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type VisitRole struct {
	infra.Model      `model:"-"` //
	RoleId           string      `model:"pk"  db:"char(36)"` //角色Id
	Act              int         `db:"int(11)"`              //角色当前行动力
	ConsumeAct       int         `db:"int(11)"`              //触发委托更新累计（<5）
	ActTime          int64       `db:"bigint(64)"`           //行动力刷新时间
	LastDispatchTime int64       `db:"bigint(20)"`           //最后派遣时间
	UnlockedCountry  int         `db:"int(11)"`              //已经解锁的国家，用来提示新解锁国家
	DispatchTimes    string      `db:"varchar(255)"`         //派遣次数
	RecoveryTimes    string      `db:"varchar(255)"`         //恢复次数
	MainStory        int         `db:"int(11)"`              //主线剧情进度
	PlayStory        int         `db:"int(11)"`              //是否播放剧情对话（0.否 1.播放）
	StoryOver        int         `db:"int(11)"`              //是否为最后一场剧情（0.否  1.是）
	VisitTime        int64       `db:"bigint(20)"`           //最后更新寻访的时间
}
