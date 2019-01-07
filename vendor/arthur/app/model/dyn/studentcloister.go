/*
Created on 2018-11-06 16:58:42
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type StudentCloister struct {
	infra.Model  `model:"-"`                        //
	RoleId       string `model:"pk"  db:"char(36)"` //玩家id
	Level        int16  `db:"smallint(5)"`          //修道院等级
	Exp          int    `db:"int(10)"`              //修道院经验
	RecoverTime  int64  `db:"bigint(20)"`           //鼓励次数恢复至满的时间
	LastGrowTime int64  `db:"bigint(20)"`           //上次成长时间
}
