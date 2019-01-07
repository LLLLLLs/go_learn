/*
Created on 2018-07-09 16:35:40
author: Auto Generate
*/
package stat

type HeroStar struct {
	Star    int16 `model:"pk"` //星级
	Prop    int16 //升级概率百分数
	NeedExp int   //消耗资质经验
}
