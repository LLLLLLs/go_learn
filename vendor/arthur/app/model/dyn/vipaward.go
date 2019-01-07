/*
Created on 2018-11-19 16:17:44
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type VipAward struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	RoleId      string `db:"char(36)"`             //
	Level       int16  `db:"smallint(5)"`          //vip奖励等级
	Time        int64  `db:"bigint(20)"`           //领取时间
}
