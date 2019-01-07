/*
Created on 2018-08-15 16:07:35
author: Auto Generate
*/
package stat

import "arthur/app/enum/AwardType"

type Award struct {
	GroupId string         //奖励所属群组id
	Typ     AwardType.Type //奖励类型: 值奖励，道具奖励
	No      int16          //奖励编号
	Count   int64          //奖励数量
	Weight  int            //奖励在群组中的权重
}
