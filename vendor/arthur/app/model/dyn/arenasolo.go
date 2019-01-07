/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ArenaSolo struct {
	infra.Model  `model:"-"`                        //
	Id           string `model:"pk"  db:"char(36)"` //1V1战斗记录id
	RoleId       string `db:"char(36)"`             //角色id
	AppointId    string `db:"char(36)"`             //指定派遣的记录ID(挑战/复仇)
	EnemyId      string `db:"char(36)"`             //对手id
	Typ          int16  `db:"smallint(5)"`          //1=随机 2=道具随机 3=挑战 4=复仇
	IsNpc        bool   `db:"tinyint(1)"`           //是否为NPC(npc战斗结束后不生成防守记录)
	HeroNo       int16  `db:"smallint(5)"`          //英雄编号
	SelfHeroLv   int16  `db:"smallint(5)"`          //我方英雄等级
	EnemyHeroLv  int16  `db:"smallint(5)"`          //敌方英雄等级
	SelfAttack   int64  `db:"bigint(20)"`           //我方攻击
	SelfBlood    int64  `db:"bigint(20)"`           //我方血量
	EnemyAttack  int64  `db:"bigint(20)"`           //敌方攻击
	EnemyBlood   int64  `db:"bigint(20)"`           //敌方血量
	GenerateTime int64  `db:"bigint(20)"`           //记录生成时间
	CompleteTime int64  `db:"bigint(20)"`           //战斗结束时间
	Skill1       int16  `db:"smallint(5)"`          //技能1
	Skill2       int16  `db:"smallint(5)"`          //技能2
	Skill3       int16  `db:"smallint(5)"`          //技能3
	RandSeed     int    `db:"int(10)"`              //随机种子
	RefreshTimes int16  `db:"smallint(5)"`          //技能刷新次数
	Result       int16  `db:"smallint(5)"`          //战斗结果-->用于竞技名录筛选 0=无 1=己方上榜 2=敌方上榜
	Accepted     bool   `db:"tinyint(1)"`           //是否接受战斗
}
