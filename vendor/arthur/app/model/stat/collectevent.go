/*
Created on 2018-09-07 15:22:57
author: Auto Generate
*/
package stat

type CollectEvent struct {
	Id   int16  `model:"pk"` //事件ID
	Text string //事件内容
	Type int16  //事件类型 1=银两，2=粮食，3=士兵
}
