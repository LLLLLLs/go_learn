/*
Created on 2018-11-16 10:10:00
author: Auto Generate
*/
package stat

type MainTask struct {
	Id              int16  `model:"pk"  db:"smallint(6)"` //主键
	Title           string `db:"char(32)"`                //主线标题
	TaskType        int16  `db:"smallint(6)"`             //任务类型
	ValueNo         int16  `db:"smallint(6)"`             //关注事件编号
	RequireProgress int16  `db:"smallint(6)"`             //要求进度
	SceneNo         int16  `db:"smallint(6)"`             //场景编号
	SpecTarget      int16  `db:"smallint(6)"`             //特定目标
	IsDelta         int16  `db:"smallint(6)"`             //是否关注增量，1为关注，0为关注全量
	AwardGroup      string `db:"char(36)"`                //完成任务时奖励组ID
}
