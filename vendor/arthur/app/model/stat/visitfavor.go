/*
Created on 2018-11-05 11:41:26
author: Auto Generate
*/
package stat

type VisitFavor struct {
	Id       int     `model:"pk"` //
	IncFavor float64 //增加好感度概率
	Unlock   float64 //解锁名媛概率
	Meet     float64 //保持原有好感度
	BeautyId int     //名媛编号
	Favor    int     //好感度
}
