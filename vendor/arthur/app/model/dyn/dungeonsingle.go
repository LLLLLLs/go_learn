/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type DungeonSingle struct {
	infra.Model `model:"-"` //
	RoleId      string      `model:"pk"  db:"char(36)"` //角色id
	DragonBlood int64       `db:"bigint(20)"`           //巨龙当前血量
	Hero1       int16       `db:"smallint(5)"`          //英雄1
	Hero2       int16       `db:"smallint(5)"`          //英雄1血量
	Hero3       int16       `db:"smallint(5)"`          //英雄2
	Hero4       int16       `db:"smallint(5)"`          //英雄2血量
	Blood1      int64       `db:"bigint(20)"`           //英雄3
	Blood2      int64       `db:"bigint(20)"`           //英雄3血量
	Blood3      int64       `db:"bigint(20)"`           //英雄4
	Blood4      int64       `db:"bigint(20)"`           //英雄4血量
}
