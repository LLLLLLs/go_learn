/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type MonthCard struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	RoleId      string `db:"char(36)"`             //
	Type        int16  `db:"smallint(5)"`          //月卡类型 1=基础月卡 2=高级月卡 3=年卡
	StartTime   int64  `db:"bigint(20)"`           //激活时间
	AwardTime   int64  `db:"bigint(20)"`           //奖励领取时间
}
