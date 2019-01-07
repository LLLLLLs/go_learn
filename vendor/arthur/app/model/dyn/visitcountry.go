/*
Created on 2018-11-06 16:58:42
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type VisitCountry struct {
	infra.Model     `model:"-"` //
	Id              string      `model:"pk"  db:"char(36)"` //
	RoleId          string      `db:"char(36)"`             //
	CountryId       int         `db:"int(11)"`              //
	Lv              int         `db:"int(255)"`             //国家等级
	Exp             int64       `db:"bigint(20)"`           //国家经验
	EntrustEventNum int         `db:"int(11)"`              //委托事件数量，上限10
	MainEventId     int         `db:"int(11)"`              //主线剧情
	Dlv             int         `db:"int(11)"`              //外交等级
	Dexp            int         `db:"int(11)"`              //外交经验
}
