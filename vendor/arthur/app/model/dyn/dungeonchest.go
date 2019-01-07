/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type DungeonChest struct {
	infra.Model `model:"-"`
	RoleId      string `model:"pk"  db:"char(36)"` //
	Time        int64  `db:"bigint(20)"`           //最新购买时间
	SilverChest int16  `db:"smallint(5)"`          //白银宝箱购买次数
	GoldenChest int16  `db:"smallint(5)"`          //黄金宝箱购买次数
}
