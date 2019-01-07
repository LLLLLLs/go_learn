/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ArenaHeroRound struct {
	infra.Model    `model:"-"`
	Id             string `model:"pk"  db:"char(36)"` //
	FightId        string `db:"char(36)"`             //本场战斗id
	RoundId        string `db:"char(36)"`             //本轮战斗id
	PositionNo     int16  `db:"smallint(6)"`          //英雄站位编号
	RoleId         string `db:"char(36)"`             //角色id
	HeroNo         int16  `db:"smallint(6)"`          //英雄编号
	HeroBlood      int64  `db:"bigint(20)"`           //英雄血量
	HeroBloodLimit int64  `db:"bigint(20)"`           //英雄血量上限
	HeroAttack     int64  `db:"bigint(20)"`           //英雄攻击力
	IsAttacked     int16  `db:"smallint(4)"`          //0为攻方，1为守方
	Status         int16  `db:"smallint(4)"`          //本场战斗状态：0未结束，1已结束
	RoundWin       int16  `db:"smallint(4)"`          //本轮胜负：0未确定；1胜利；2失败
	MatchTime      int64  `db:"bigint(20)"`           //本场战斗开始时间；匹配arena_fight表的match_time
}
