/*
Created on 2018-08-28 15:08:52
author: Auto Generate
*/
package stat

type Miracle struct {
	Type   int16 //神迹类型：1、收税；2、学习；3、骑士升级；4、膜拜；5、名媛邀约
	Chance int16 //触发概率
	Value  int   //触发数值(倍数、升级数、钻石)
}
