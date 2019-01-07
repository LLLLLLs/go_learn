/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type DungeonSingleExplore struct {
	infra.Model `model:"-"` //
	Id          string      `model:"pk"  db:"char(36)"` //
	RoleId      string      `db:"char(36)"`             //角色id
	Choice      int16       `db:"smallint(5)"`          //探索地点
}
