/*
Created on 2018-11-13 17:44:38
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type WishChoice struct {
	infra.Model `model:"-"`
	RoleId      string `model:"pk"  db:"char(36)"` //玩家id
	Choice1     int16  `db:"smallint(5)"`          //选项1
	Choice2     int16  `db:"smallint(5)"`          //选项2
	Choice3     int16  `db:"smallint(5)"`          //选项3
	Choice4     int16  `db:"smallint(5)"`          //选项4
	RefreshTime int64  `db:"bigint(20)"`           //刷新时间
}
