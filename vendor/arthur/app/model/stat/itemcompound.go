/*
Created on 2018-08-21 15:03:38
author: Auto Generate
*/
package stat

import (
	"arthur/app/enum/AwardType"
	"arthur/app/enum/ItemNo"
)

type ItemCompound struct {
	ItemNo    ItemNo.Type    //合成目标道具编号
	NeedTyp   AwardType.Type //需求的类型：1值，2道具
	NeedNo    int16          //需求编号
	NeedCount int64          //需求数量
}
