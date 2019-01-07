/*
Created on 2018-11-06 16:58:42
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ResBonus struct {
	infra.Model  `model:"-"`
	Id           string `model:"pk"  db:"char(36)"` //事件id
	RoleId       string `db:"char(36)"`             //角色id
	Typ          int16  `db:"smallint(2)"`          //事件类型
	HeroNo       int16  `db:"smallint(5)"`          //执行事件的英雄编号
	TextId       int16  `db:"smallint(5)"`          //事件文本id
	GenerateTime int64  `db:"bigint(20)"`           //事件触发时间
	Completed    bool   `db:"tinyint(4)"`           //事件是否完成
}
