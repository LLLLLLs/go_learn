/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type DailyActivity struct {
	infra.Model  `model:"-"`
	Id           string `model:"pk"  db:"varchar(64)"` //
	RoleId       string `db:"char(64)"`                //
	ActivityId   int16  `db:"smallint(11)"`            //活跃度奖励档位
	GetAwardTime int64  `db:"bigint(255)"`             //领取时间
}
