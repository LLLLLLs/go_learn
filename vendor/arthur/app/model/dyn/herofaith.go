/*
Created on 2018-12-25 15:54:45
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type HeroFaith struct {
	infra.Model     `model:"-"`
	Id              string `model:"pk"  db:"char(36)"` //
	HeroId          string `db:"char(36)"`             //
	FaithSkillNo    int16  `db:"smallint(5)"`          //信仰技能编号
	FaithSkillLevel int16  `db:"smallint(5)"`          //信仰技能等级
}
