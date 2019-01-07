/*
Created on 2018-10-22 17:10:47
author: Auto Generate
*/
package stat

type VisitCountryNpcEvent struct {
	Id       int      `model:"pk"` //
	Award    [][2]int //奖励列表
	NpcStage int      //大事记阶段
	NpcId    int      //npc
	EventId  int      //对应事件编号
}
