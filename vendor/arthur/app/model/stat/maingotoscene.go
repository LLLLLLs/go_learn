/*
Created on 2018-11-16 09:11:49
author: Auto Generate
*/
package stat

type MainGotoScene struct {
	SceneNo   int16  `model:"pk"  db:"smallint(6)"` //主线任务跳转场景编号
	SceneDesc string `db:"char(32)"`                //
}
