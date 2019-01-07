/*
Created on 2018-09-20 15:20:34
author: Auto Generate
*/
package stat

type AchievementDesc struct {
	No       int16  `model:"pk"` //成就编号Id
	Name     string //成就名称
	Describe string //成就描述
}
