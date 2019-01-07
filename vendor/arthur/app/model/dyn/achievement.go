/*
Created on 2018-11-06 16:58:39
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type Achievement struct {
	infra.Model   `model:"-"`
	Id            string `model:"pk"  db:"char(36)"` //
	RoleId        string `db:"char(64)"`             //角色id
	AchievementId int16  `db:"smallint(5)"`          //成就id
	CurrentNum    int64  `db:"bigint(64)"`           //当前实现数目
	GetStageId    int16  `db:"smallint(5)"`          //已领取的奖励档位
	Finish        int16  `db:"smallint(5)"`          //达成该系全部成就 (1.已达成 0.未达成)
}
