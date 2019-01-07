/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ArenaReloadHeroRecord struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	RoleId      string `db:"char(36)"`             //角色ID
	RoundId     string `db:"char(36)"`             //本轮战斗ID
	FightId     string `db:"char(36)"`             //本场战斗ID
	RoundWin    int16  `db:"smallint(6)"`          //本轮胜负：0未确定；1胜利；2失败
	Status      int16  `db:"smallint(6)"`          //本场战斗状态：0未结束，1已结束
	ReloadTime  int64  `db:"bigint(20)"`           //替换时间
}
