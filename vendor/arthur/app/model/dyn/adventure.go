/*
Created on 2018-11-06 16:58:39
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type Adventure struct {
	infra.Model    `model:"-"` //
	RoleId         string      `model:"pk"  db:"char(36)"` //
	LastEnterCity  int64       `db:"bigint(20)"`           //最后一次进入城堡时间
	AdventureTimes int         `db:"int(11)"`              //今日已经使用探险次数
	VisitedPos     int64       `db:"bigint(11)"`           //当前城堡已探险地点id的总和
	CurrentCity    int         `db:"int(11)"`              //当前城堡
	IsSuccess      bool        `db:"tinyint(4)"`           //是否打败boss
	TodayHero      int         `db:"int(11)"`              //当前英雄
	Seed           int         `db:"int(11)"`              //seed
	AwardPos       int64       `db:"bigint(20)"`           //已经获取奖励的地点
	RefreshSkill   int         `db:"int(11)"`              //boss技能刷新次数
	Dispatch       string      `db:"varchar(255)"`         //派遣次数
	Recovery       string      `db:"varchar(255)"`         //恢复派遣次数
	BeatBoss       string      `db:"varchar(255)"`         //打败boss的次数
	ResetTimes     int         `db:"int(11)"`              //重置最后一个城堡的次数
	Hint           int64       `db:"bigint(255)"`          //已经提示的城堡
}
