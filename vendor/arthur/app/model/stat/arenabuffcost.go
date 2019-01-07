/*
Created on 2018-10-19 16:32:12
author: Auto Generate
*/
package stat

type ArenaBuffCost struct {
	No    int16 `model:"pk"` //buff消耗主键，no同时代表次数
	Pearl int16 //所需竞技宝珠，宝珠不足消耗元宝
	Gold  int16 //所需钻石数量
}
