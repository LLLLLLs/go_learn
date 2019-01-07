/*
Created on 2018-09-18 16:56:45
author: Auto Generate
*/
package stat

type DailyActivity struct {
	ActivityId      int16  `model:"pk"` //活跃度奖励档位
	RequireActivity int64  //活跃度要求
	GroupId         string //奖励组
	AwardRate       int64  //档位奖励倍率
	BaseNum         int64  //基础奖励数值
}
