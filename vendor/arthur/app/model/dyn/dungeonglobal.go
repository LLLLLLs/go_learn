/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type DungeonGlobal struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	MaxBlood    int64  `db:"bigint(20)"`           //最大血量
	KillerId    string `db:"char(36)"`             //击杀者id
	StartTime   int64  `db:"bigint(20)"`           //活动开始时间（活动当天的零点）
	IsClosed    bool   `db:"tinyint(1)"`           //该活动是否被关闭
}
