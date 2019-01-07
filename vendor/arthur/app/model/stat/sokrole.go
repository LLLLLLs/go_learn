/*
Created on 2018-07-11 16:16:27
author: Auto Generate
*/
package stat

type SokRole struct {
	No      int     `model:"pk"` // 王者之剑编号
	Typ     int     // 加成类型:1、商;2、农;3、政;4、军
	Num     int64   // 加成数值
	Percent float64 // 加成百分比
}
