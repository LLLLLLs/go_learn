/*
Created on 2018-07-11 11:37:09
author: Auto Generate
*/
package stat

type SokBeauty struct {
	No       int     `model:"pk"` // 王者之剑编号
	Affected []int   // 受影响的红颜编号,空代表全部
	Intimacy int     // 亲密度
	Charm    int     // 魅力值
	Exp      float64 // 经验加成百分比
}
