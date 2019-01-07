/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import (
	"arthur/app/enum/MiracleType"
	"arthur/app/infra"
)

type Miracle struct {
	infra.Model `model:"-"`      //
	Id          string           `model:"pk"  db:"char(36)"` //
	RoleId      string           `db:"char(36)"`             //角色id
	Type        MiracleType.Type `db:"smallint(5)"`          //神迹类型：1、收税；2、学习；3、骑士升级；4、膜拜；5、名媛邀约
	Times       int16            `db:"smallint(5)"`          //该类型神迹触发次数
	RefreshTime int64            `db:"bigint(20)"`           //神迹触发时间
}
