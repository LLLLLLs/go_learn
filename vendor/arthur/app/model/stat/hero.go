/*
Created on 2018-07-10 09:17:07
author: Auto Generate
*/
package stat

import (
	"arthur/app/enum/AttrType"
	"arthur/app/enum/HeroFaithType"
	"arthur/app/enum/HeroProfNo"
)

type Hero struct {
	No       int16              `model:"pk"  db:"smallint(5)"` //编号
	Name     string             `db:"char(36)"`                //名字
	Race     int16              `db:"smallint(5)"`             //种族
	FaithTyp HeroFaithType.Type `db:"smallint(5)"`             //信仰类型，0： 无信仰，1： 2： 3：
	Talent   []int16            `db:"char(128)"`               //资质编号列表
	Spec     []AttrType.Type    `db:"char(50)"`                //特长，值为属性类型， 智慧：1, 勤勉：2, 忠诚：3, 英勇：4
	ProfNo   HeroProfNo.Type    `db:"smallint(4)"`             //职业编号
	IsFree   bool               `db:"tinyint(4)"`              //是否免费
}
