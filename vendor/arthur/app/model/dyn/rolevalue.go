/*
Created on 2018-11-06 16:58:42
author: Auto Generate
*/
package dyn

import (
	"arthur/app/enum/RoleValueNo"
	"arthur/app/infra"
	"encoding/gob"
)

type RoleValue struct {
	infra.Model `model:"-"`
	Id          string           `model:"pk"  db:"char(36)"` //
	RoleId      string           `db:"char(36)"`             //
	ValNo       RoleValueNo.Type `db:"smallint(5)"`          //值类型
	Val         int64            `db:"bigint(20)"`           //数值
	RefreshTime int64            `db:"bigint(20)"`           //数值更新时间
	Description string           `db:"varchar(255)"`         //描述，开发过程中便于debug
}

func init() {
	gob.Register(RoleValueNo.Lv)
}
