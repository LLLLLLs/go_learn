/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ArenaChallengeRecord struct {
	infra.Model   `model:"-"`
	Id            string `model:"pk"  db:"char(36)"` //
	RoleId        string `db:"char(36)"`             //角色ID
	RecordId      string `db:"char(36)"`             //战斗ID，来自Arena_fight表主键
	ChallengeType int16  `db:"smallint(6)"`          //开战类型，0为竞技名录挑战，1为防守记录复仇
	Status        int16  `db:"smallint(6)"`          //本场战斗状态，0未结束，1已结束
	ChallengeTime int64  `db:"bigint(20)"`           //挑战时间
}
