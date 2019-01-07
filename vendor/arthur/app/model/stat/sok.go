/*
Created on 2018-07-11 16:16:21
author: Auto Generate
*/
package stat

type Sok struct {
	No    int `model:"pk"` // 王者之剑编号
	Typ   int // 加成类型：1、属性；2、征收；3、子嗣；4、红颜
	Level int // 解锁关卡数
}
