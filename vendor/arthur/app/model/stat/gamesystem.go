/*
Created on 2018-11-22 15:34:15
author: Auto Generate
*/
package stat

import "arthur/app/enum/ModuleNo"

type GameSystem struct {
	No             ModuleNo.Type  `model:"pk"  db:"int(5)"` //模块编号
	Name           string `db:"varchar(255)"`       //名称
	Rule           string `db:"varchar(255)"`       //规则
	ParentNo       ModuleNo.Type  `db:"int(5)"`             //父模块编号
	UnlockNoList   []int16 `db:"varchar(255)"`		//解锁要求
}
