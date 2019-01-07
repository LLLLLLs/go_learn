/*
Created on 2018-09-28 10:09:52
author: Auto Generate
*/
package stat

import "arthur/app/enum/AttrType"

type BeautySkills struct {
	SkillId       int     `model:"pk"` //
	Name          string  //红颜技能名称
	AmityRequired int     //解锁技能需要的友好度
	AddType       []AttrType.Type  //加成属性的类型，1武力，2智力，3政治，4魅力，可在列表中增加多项
	AddValue      int     //红颜加成具体值
	AddPercent    float64 //红颜加成具体值百分比, add_value不为0时，以add_value为准
}
