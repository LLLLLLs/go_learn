/*
Created on 2018-11-08 11:03:50
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type VisitEvent struct {
	infra.Model    `model:"-"` //
	Id             string      `model:"pk"  db:"char(36)"` //
	RoleId         string      `db:"char(36)"`             //role id
	EntrustLv      int         `db:"int(11)"`              //事件等级
	EntrustId      int         `db:"int(20)"`              //当前委托事件id
	CountryId      int         `db:"int(11)"`              //国家
	VisitId        int         `db:"int(11)"`              //寻访事件id
	EntrustCity    int         `db:"int(255)"`             //委托城市位置
	VisitCity      int         `db:"int(255)"`             //寻访城市位置
	LastEntrustId  int         `db:"int(20)"`              //最后的线性委托事件id
	EntrustCreated int64       `db:"bigint(255)"`          //委托事件创建时间
	VisitCreated   int64       `db:"bigint(255)"`          //寻访事件创建时间
}
