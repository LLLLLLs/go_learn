/*
Created on 2018-11-06 16:58:42
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type Wish struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	RoleId      string `db:"char(36)"`             //角色id
	Day         int16  `db:"smallint(5)"`          //许愿天数
	Choice      int16  `db:"smallint(5)"`          //许愿选项
	Total       bool   `db:"tinyint(1)"`           //是否为累计奖励
}
