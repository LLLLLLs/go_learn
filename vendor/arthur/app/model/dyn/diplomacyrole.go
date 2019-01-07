/*
Created on 2018-11-09 11:58:19
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type DiplomacyRole struct {
	infra.Model     `model:"-"` //
	RoleId          string      `model:"pk"  db:"char(36)"` //
	VictoryNum      int         `db:"int(11)"`              //连胜次数
	RiotNum         int         `db:"int(11)"`              //暴动次数
	DictatorshipNum int         `db:"int(11)"`              //独裁次数
	UpdateTime      int64       `db:"bigint(32)"`           //议案处理次数（act）刷新时间
	PowerA          int         `db:"int(11)"`              //教会势力值
	PowerB          int         `db:"int(11)"`              //军队势力值
	PowerC          int         `db:"int(11)"`              //商会势力值
	PowerD          int         `db:"int(11)"`              //平民势力值
	Act             int         `db:"int(11)"`              //议政次数
	RefreshTime     int64       `db:"bigint(32)"`           //上一次议案刷新时间
	MaxVictoryNum   int         `db:"int(11)"`              //最大连胜数
}
