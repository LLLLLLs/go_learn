/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ArenaPerfectRecord struct {
	infra.Model   `model:"-"`
	Id            string `model:"pk"  db:"char(36)"` //全歼记录主键来自战斗记录主键
	AttackRoleId  string `db:"char(36)"`             //攻方角色ID
	AttackHerosNo string `db:"varchar(32)"`          //攻方英雄编号
	AttackNum     int16  `db:"smallint(6)"`          //击杀次数
	EnemyRoleId   string `db:"char(36)"`             //敌方角色id
	IsNpc         int16  `db:"smallint(6)"`          //敌方角色是否是NPC，0为是，1为否
	RecordTime    int64  `db:"bigint(20)"`           //上榜时间
}
