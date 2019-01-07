/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ArenaFightRound struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	RoleId      string `db:"char(36)"`             //角色ID
	RoundNum    int16  `db:"smallint(6)"`          //第几轮
	FightId     string `db:"char(36)"`             //战斗外键
	Win         int16  `db:"smallint(4)"`          //本轮胜负状态：0尚未确定；1胜利；2失败
	Status      int16  `db:"smallint(4)"`          //本场战斗状态：0为尚未结束，1为已结束
	StartTime   int64  `db:"bigint(20)"`           //开本轮始战斗时间
}
