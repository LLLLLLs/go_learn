/*
Created on 2018-10-30 17:06:01
author: Auto Generate
*/
package stat

type VisitCountry struct {
	Id            int    `model:"pk"` //国家id
	Name          string //
	CityList      []int  //所含城市列表id
	HeroList      []int  //所含英雄列表
	BeautyList    []int  //名媛列表
	CountryTalent []int  //国家特长
	NpcList       []int  //npc列表
}
