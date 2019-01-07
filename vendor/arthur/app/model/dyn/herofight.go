/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import (
	"arthur/app/enum/HeroFightType"
	"arthur/app/infra"
)

type HeroFight struct {
	infra.Model  `model:"-"`
	Id           string             `model:"pk"  db:"char(36)"` //
	HeroNo       int16              `db:"smallint(5)"`          //英雄编号
	RoleId       string             `db:"char(36)"`             //角色id
	FightTimes   int16              `db:"smallint(5)"`          //已战斗次数
	RecoverTimes int16              `db:"smallint(5)"`          //已回复次数
	FightTime    int64              `db:"bigint(20)"`           //战斗时间
	Typ          HeroFightType.Type `db:"smallint(5)"`          //战斗类型
}
