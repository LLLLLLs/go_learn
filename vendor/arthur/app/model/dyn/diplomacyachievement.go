/*
Created on 2018-11-19 14:40:17
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type DiplomacyAchievement struct {
	infra.Model  `model:"-"` //
	Id           string      `model:"pk"  db:"char(32)"` //
	RoleId       string      `db:"char(32)"`             //
	No           int         `db:"int(11)"`              //外交成就编号
	IsReceive    int16       `db:"smallint(6)"`          //是否领取
	GetAwardTime int64       `db:"bigint(32)"`           //领取时间
	ReceiveTime  int         `db:"int(32)"`              //领取次数
}
