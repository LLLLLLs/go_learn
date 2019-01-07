/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ArenaFightMemory struct {
	infra.Model  `model:"-"`
	Id           string `model:"pk"  db:"char(36)"` //战斗记忆主键
	AttackRoleId string `db:"char(36)"`             //攻方角色ID
	EnemyRoleId  string `db:"char(36)"`             //守方角色ID
	IsNpc        int16  `db:"smallint(6)"`          //是否为与NPC战斗，0是，1不是
	FightType    int16  `db:"smallint(6)"`          //开战类型，0为免费；1为随机（出使令）
	MatchTime    int64  //战斗匹配时间
}
