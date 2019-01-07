/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ArenaSoloDefence struct {
	infra.Model  `model:"-"`
	Id           string `model:"pk"  db:"char(36)"` //
	IsRevenged   bool   `db:"tinyint(1)"`           //是否复仇
	AttackId     string `db:"char(36)"`             //进攻方id
	DefenseId    string `db:"char(36)"`             //防守方id
	HeroNo       int16  `db:"smallint(5)"`          //英雄编号
	GenerateTime int64  `db:"bigint(20)"`           //流水创建时间
	Result       int16  `db:"smallint(5)"`          //0=失败 1=险胜 2=小胜 3=大胜 4=完胜
}
