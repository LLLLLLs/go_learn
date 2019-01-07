/*
Created on 2018-09-06 14:06:48
author: Auto Generate
*/
package stat

type BeautyEventAward struct {
	Id          int     `model:"pk"` //场景id
	No          int     //item表中的道具id
	Number      int     //道具数量
	AwardType   int     //奖励类型， 1=小奖励， 其它=大奖励
	Probability float64 //获取概率,等概率填1
}
