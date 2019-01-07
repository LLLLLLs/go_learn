/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ArenaStorePurchase struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	RoleId      string `db:"char(36)"`             //玩家id
	ItemNo      int16  `db:"smallint(5)"`          //商品编号
	Times       int16  `db:"smallint(5)"`          //购买次数
	UpdateTime  int64  `db:"bigint(20)"`           //购买记录更新时间
}
