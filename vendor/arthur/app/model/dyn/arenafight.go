/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ArenaFight struct {
	infra.Model  `model:"-"`
	Id           string `model:"pk"  db:"char(36)"` //主键
	AttackRoleId string `db:"char(36)"`             //玩家角色id
	EnemyRoleId  string `db:"char(36)"`             //敌方角色id
	IsNpc        int16  `db:"smallint(6)"`          //是否为与NPC战斗，0是，1不是
	Status       int16  `db:"smallint(4)"`          //0为尚未开战；1接受战斗；2至少开始一轮战斗；3战斗结束
	Win          int16  `db:"smallint(4)"`          //0为未确定；1为胜利；2失败
	FightType    int16  `db:"smallint(4)"`          //开战类型，0为免费；1为随机（出使令）；2为挑战书
	MatchTime    int64  `db:"bigint(20)"`           //匹配到对手时间
	ReceiveTime  int64  `db:"bigint(20)"`           //接受时间
	EndTime      int64  `db:"bigint(20)"`           //一场战斗结束时间
}
