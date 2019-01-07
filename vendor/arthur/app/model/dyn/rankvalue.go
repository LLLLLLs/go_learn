/*
Created on 2018-11-06 16:58:42
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type RankValue struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	RankId      string `db:"char(36)"`             //排行榜id
	TargetId    string `db:"char(36)"`             //数值目标id
	Value       int64  `db:"bigint(20)"`           //数值
	UpdateDate  int64  `db:"bigint(20)"`           //更新时间
}
