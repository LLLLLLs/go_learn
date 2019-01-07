/*
Created on 2018-09-18 09:19:47
author: Auto Generate
*/
package stat

type DailyTask struct {
	No         int16  `model:"pk"` //任务类型
	RequireNum int64  //获得奖励所需数值
	Liveness   int64  //奖励活跃度
	GroupId    string //奖励group_id
	Desc       string //说明
	Jump       int16  //客户端界面跳转
	Exp        int64  //经验值奖励
}
