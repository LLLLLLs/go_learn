/*
Created on 2018-10-16 14:14:03
author: Auto Generate
*/
package stat

type MonthCardAward struct {
	Type          int16 `model:"pk"` //
	AwardOnce     string             //一次性奖励
	AwardEveryday string             //每日奖励
}
