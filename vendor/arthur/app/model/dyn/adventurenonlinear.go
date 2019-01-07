/*
Created on 2018-11-26 17:28:57
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type AdventureNonlinear struct {
	infra.Model  `model:"-"`
	Id           string `model:"pk"  db:"char(36)"` //
	RoleId       string `db:"char(36)"`             //
	CityId       int    `db:"int(11)"`              //非线性城堡id
	TodayHero    int    `db:"int(255)"`             //今日派遣英雄
	VisitedPos   int64  `db:"bigint(255)"`          //已探险地点
	AwardPos     int64  `db:"bigint(255)"`          //已经获取奖励的地点
	IsSuccess    bool   `db:"tinyint(4)"`           //是否打败boss
	RefreshSkill int    `db:"int(11)"`              //boss技能刷新次数
	IsHint       bool   `db:"tinyint(255)"`         //已经提示解锁的城堡
	Seed         int    `db:"int(11)"`              //刷新技能需要的seed
}
