/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type HeroTalent struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	No          int16  `db:"smallint(4)"`          //
	Lv          int16  `db:"smallint(4)"`          //天赋等级
	HeroId      string `db:"char(36)"`             //英雄id
}
