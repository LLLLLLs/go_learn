/*
Created on 2018-10-18 11:27:35
author: Auto Generate
*/
package stat

type VisitEvent struct {
	Id          int16   `model:"pk"` //寻访事件
	Typ         int16   //寻访类型(1.贸易、2.游玩、3.巡视、 0.任意)
	GroupId     string  //事件奖励基础值分组
	Coefficient float64 //奖励系数
	Description string  //事件描述
	Probability int16   //发生概率
	AwardText   string  //奖励文本
}
