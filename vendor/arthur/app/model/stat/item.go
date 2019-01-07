/*
Created on 2018-08-22 13:50:48
author: Auto Generate
*/
package stat

import (
	"arthur/app/enum/ItemNo"
	"arthur/app/enum/ItemType"
)

type Item struct {
	No         ItemNo.Type   `model:"pk"` //
	Name       string        //道具名
	Typ        ItemType.Type //类型，根据效果分类
	EffectData []byte        //使用效果数据
}
