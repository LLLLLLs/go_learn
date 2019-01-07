/*
Created on 2018-11-16 10:10:15
author: Auto Generate
*/
package stat

type MainTaskType struct {
	TaskType int16  `model:"pk"  db:"smallint(6)"` //主线任务类型
	TypeDesc string `db:"char(32)"`                //
}
