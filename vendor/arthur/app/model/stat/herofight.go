/*
Created on 2018-10-24 15:55:53
author: Auto Generate
*/
package stat

import "arthur/app/enum/HeroFightType"

type HeroFight struct {
	Typ             HeroFightType.Type //战斗类型：1远征，2单人副本，3全局副本，4,竞技场1v1随机，5竞技场制定派遣次数
	CanFreeTimes    int16 //可免费出站次数
	CanRestoreTimes int16 //可付费出站次数
}
