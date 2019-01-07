/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ArenaSoloChallenge struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	RoleId      string `db:"char(36)"`             //角色id
	FlowId      string `db:"char(36)"`             //挑战的流水ID
	Time        int64  `db:"bigint(20)"`           //挑战时间
}
