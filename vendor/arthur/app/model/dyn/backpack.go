/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type Backpack struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	ItemNo      int16  `db:"smallint(4)"`          //
	RoleId      string `db:"char(36)"`             //
	Count       int64  `db:"bigint(20)"`           //拥有道具数量
}
