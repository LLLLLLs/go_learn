/*
Created on 2018-09-10 17:25:13
author: Auto Generate
*/
package stat

import (
	"arthur/app/enum/RoleValueNo"
	"arthur/app/enum/ValueEventNo"
)

type RoleValue struct {
	No           RoleValueNo.Type  `model:"pk"` //值编号
	ValueEventNo ValueEventNo.Type //对应的值事件
	Description  string            //描述
	CanNegative  bool              //是否可为负值
}
