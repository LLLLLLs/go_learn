/*
Created on 2018-12-25 15:54:31
author: Auto Generate
*/
package stat

import "arthur/app/enum/HeroFaithType"

type HeroFaithSkill struct {
	Id             int16                   `model:"pk"  db:"smallint(5)"` //技能编号
	Name           string                  `db:"varchar(36)"`             //技能名称
	Typ            HeroFaithType.SkillType `db:"smallint(5)"`             //技能加成类型：1=等级加成；2=同信仰英雄数量加成
	WiseEffect     float64                 `db:"varchar(10)"`             //智慧加成
	DiligentEffect float64                 `db:"varchar(10)"`             //勤勉加成
	LoyaltyEffect  float64                 `db:"varchar(10)"`             //忠诚加成
	HeroicEffect   float64                 `db:"varchar(10)"`             //英勇加成
	LevelLimit     int16                   `db:"smallint(5)"`             //技能等级上限
	NeedItemNo     int16                   `db:"smallint(5)"`             //升级所需道具编号
	NeedItemCount  int16                   `db:"smallint(5)"`             //升级所需道具数量
	Desc           string                  `db:"varchar(255)"`            //技能描述
}
