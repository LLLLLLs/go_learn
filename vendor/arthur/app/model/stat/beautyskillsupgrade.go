/*
Created on 2018-07-23 17:57:46
author: Auto Generate
*/
package stat

type BeautySkillsUpgrade struct {
	SkillLv int16 `model:"pk"` //技能等级
	SkillA  int64 //技能1升级至对应等级所需的经验
	SkillB  int64 //
	SkillC  int64 //
	SkillD  int64 //
	SkillE  int64 //
	SkillF  int64 //
}
