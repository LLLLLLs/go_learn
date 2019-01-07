/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ChatRewardSend struct {
	infra.Model `model:"-"` //
	Id          string      `model:"pk"  db:"char(36)"` //奖励id
	SendRoleId  string      `db:"char(36)"`             //发放奖励的roleId
	Source      int16       `db:"smallint(11)"`         //来源(如1巨龙巢)
	Time        int64       `db:"bigint(36)"`           //
	MsgId       string      `db:"char(36)"`             //消息id
}
