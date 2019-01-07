/*
Created on 2018-09-21 17:54:13
author: Auto Generate
*/
package stat

import (
	"arthur/app/enum/TargetType"
	"arthur/app/enum/ValueEventNo"
)

type ValueEvent struct {
	ValueNo     ValueEventNo.Type `model:"pk"` //值编号
	TargetTyp   TargetType.Type   //目标类型
	Description string            //描述
	Name        string            //枚举名
}
