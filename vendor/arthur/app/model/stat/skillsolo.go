/*
Created on 2018-10-09 10:56:14
author: Auto Generate
*/
package stat

type SkillSolo struct {
	Id     int16  `model:"pk"` //技能id
	Typ    int16  //技能类型
	Level  int16  //技能等级
	Value1 int16  //技能效果1，百分比值：10 表示 10%
	Value2 int16  //技能效果2，百分比值：10 表示 10%
	Round  int16  //技能cd(回合数)
	Name   string //技能名字
	Desc   string //技能描述
	Weight int    //技能随机概率
}
