/*
Created on 2018-11-06 16:58:42
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type WishTotalAward struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	Day         int16  `db:"smallint(5)"`          //累计天数
	Typ         int16  `db:"smallint(5)"`          //奖励类型
	No          int16  `db:"smallint(5)"`          //奖励编号
	Count       int64  `db:"bigint(20)"`           //奖励数量
	Weight      int    `db:"int(10)"`              //所占权重
}
