/*
Created on 2018-11-06 16:58:42
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type VisitCountryNpc struct {
	infra.Model `model:"-"` //
	Id          string      `model:"pk"  db:"char(36)"` //
	NpcId       int         `db:"int(11)"`              //npc id
	RoleId      string      `db:"char(36)"`             //role id
	Stage       int         `db:"int(255)"`             //大事记阶段
	EventList   string      `db:"varchar(255)"`         //npc碰上的事件id
}
