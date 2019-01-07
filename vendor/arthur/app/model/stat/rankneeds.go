/*
Created on 2018-08-30 11:37:20
author: Auto Generate
*/
package stat

import (
	"arthur/app/enum/TargetType"
)

type RankNeeds struct {
	Id         string          `model:"pk"` //编号
	TargetType TargetType.Type //目标类型(1玩家，2英雄）
	ValueType  int16           //排行榜值类型
}
