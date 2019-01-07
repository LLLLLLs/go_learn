/*
Created on 2018-10-18 16:05:14
author: Auto Generate
*/
package stat

type VisitEntrustType struct {
	Lv                 int     `model:"pk"  db:"int(11)"` //事件等级 1-4
	Refresh            float64 `db:"float"`               //刷新概率
	EventCoefficient   float64 `db:"float"`               //事件系数
	Act                int     `db:"int(255)"`            //消耗行动力数
	BeautyLock         float64 `db:"float"`               //解锁名媛前邂逅概率
	BeautyUnlock       float64 `db:"float"`               //解锁名媛后邂逅概率
	ItemProbability    float64 `db:"float"`               //获得道具概率
	UpgradeProbability float64 `db:"float"`               //等级升级概率
	AmityProbability   float64 `db:"float"`               //解锁后获得友好度概率
}
