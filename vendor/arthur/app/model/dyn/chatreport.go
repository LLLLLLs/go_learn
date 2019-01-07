/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ChatReport struct {
	infra.Model  `model:"-"`
	Id           string `model:"pk"  db:"char(36)"` //
	RoleId       string `db:"char(36)"`             //举报者id
	BeReportedId string `db:"char(36)"`             //被举报者rid
	MsgId        string `db:"char(36)"`             //被举报消息id
	ReportTime   int64  `db:"bigint(11)"`           //举报时间
}
