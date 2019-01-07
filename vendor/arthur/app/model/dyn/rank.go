/*
Created on 2018-11-06 16:58:42
author: Auto Generate
*/
package dyn

import (
	"arthur/app/enum/RankType"
	"arthur/app/infra"
)

type Rank struct {
	infra.Model `model:"-"`
	Id          string        `model:"pk"  db:"char(36)"` //
	RankType    RankType.Type `db:"smallint(5)"`          //排行榜类型
	StartTime   int64         `db:"bigint(20)"`           //榜单开始时间
	EndTime     int64         `db:"bigint(20)"`           //榜单结束时间
}
