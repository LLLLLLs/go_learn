/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ArenaStore struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	Order       int16  `db:"smallint(5)"`          //排序
	ItemNo      int16  `db:"smallint(5)"`          //商品编号
	Price       int16  `db:"smallint(5)"`          //价格
	Limit       int16  `db:"smallint(5)"`          //限购
}
