/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type Hero struct {
	infra.Model   `model:"-"`
	Id            string `model:"pk"  db:"char(36)"` //
	No            int16  `db:"smallint(5)"`          //英雄编号
	RoleId        string `db:"char(36)"`             //角色id
	Lv            int16  `db:"smallint(4)"`          //等级
	Peerage       int16  `db:"smallint(4)"`          //爵位
	WiseAdd       int64  `db:"bigint(16)"`           //智慧道具加成
	LoyaltyAdd    int64  `db:"bigint(16)"`           //忠诚道具加成
	DiligentAdd   int64  `db:"bigint(16)"`           //勤勉加成
	HonorAdd      int64  `db:"bigint(20)"`           //
	HeroicAdd     int64  `db:"bigint(16)"`           //英勇道具加成
	TalentExp     int    `db:"int(9)"`               //资质经验
	UpgradeSilver int64  `db:"bigint(20)"`           //用于升级累计的银两
}
