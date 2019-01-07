/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type DailyTask struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	RoleId      string `db:"char(36)"`             //
	CurrentNum  int64  `db:"bigint(20)"`           //当前任务累计值（第二天重置）
	//LastUpdate   int64  `db:"bigint(20)"`           //最后更新时间
	GetAwardTime int64 `db:"bigint(20)"`  //领奖时间（第二天重置）
	TaskNo       int16 `db:"smallint(4)"` //任务编号
}
