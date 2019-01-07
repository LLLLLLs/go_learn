/*
Created on 2018-12-25 15:54:28
author: Auto Generate
*/
package stat

type HeroFaith struct {
	FaithTyp    int16   `db:"smallint(5)"` //信仰类型
	Name        string  `db:"varchar(36)"` //信仰名称
	FaithSkills []int16 `db:"varchar(36)"` //信仰技能
	HaloEffect  float64 `db:"float(10,0)"` //光环效果
	HaloLimit   int16   `db:"smallint(5)"` //光环人数加成上限
}
