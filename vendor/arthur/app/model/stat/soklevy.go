/*
Created on 2018-07-11 16:16:31
author: Auto Generate
*/
package stat

type SokLevy struct {
	No      int     `model:"pk"` // 王者之剑编号
	Typ     int     // 加成类型:1、征税;2、征粮;3、征兵
	Percent float64 // 加成百分比
}
