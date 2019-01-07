/*
Created on 2018-11-06 17:00:51
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ArenaFightRecord struct {
	infra.Model     `model:"-"`
	Id              string `model:"pk"  db:"char(36)"` //
	AttackRoleId    string `db:"char(36)"`             //攻方角色ID
	AttackHerosNo   string `db:"varchar(32)"`          //攻方出使英雄编号，用逗号分割
	AttackNum       int16  `db:"smallint(6)"`          //攻方歼灭敌方英雄数
	EnemyRoleId     string `db:"char(36)"`             //敌方角色ID
	EnemyDeltaScore int16  `db:"mediumint(9)"`         //敌方获得分数，可为负
	FightWin        int16  `db:"smallint(6)"`          //1代表全歼胜利，2代表失败
	IsNpc           int16  `db:"smallint(6)"`          //是否为与NPC战斗，0是，1不是
	RecordTime      int64  `db:"bigint(20)"`           //记录时间
}
