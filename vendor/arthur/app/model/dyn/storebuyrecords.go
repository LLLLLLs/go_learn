/*
Created on 2018-11-22 13:52:52
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type StoreBuyRecords struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //主键
	RoleId      string `db:"char(36)"`             //角色ID
	ItemId      int16  `db:"smallint(6)"`          //商城道具ID
	BuyNum      int64  `db:"bigint(20)"`           //购买数量
	BuyTime     int64  `db:"bigint(20)"`           //购买时间
}
