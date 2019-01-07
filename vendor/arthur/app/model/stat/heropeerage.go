/*
Created on 2018-07-10 10:15:24
author: Auto Generate
*/
package stat

type HeroPeerage struct {
	No          int16   `model:"pk"` //
	LvLimit     int16   //等级上限
	TalentLimit int16   //资质上限
	TalentExp   int     //资质经验奖励
	NeedItem    []int16 //升爵需要道具的id
	Color       string  //代表颜色
}
