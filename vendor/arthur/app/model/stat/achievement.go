/*
Created on 2018-09-20 15:20:29
author: Auto Generate
*/
package stat

type Achievement struct {
	No         int16  `model:"pk"` //成就编号id
	RequireNum int64  //成就要求
	Stage      int16  //成就档位
	GroupId    string //成就奖励
}
