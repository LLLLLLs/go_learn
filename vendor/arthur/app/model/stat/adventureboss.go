/*
Created on 2018-12-27 09:24:02
author: Auto Generate
*/
package stat

type AdventureBoss struct {
	Id              int     `model:"pk"  db:"int(11)"` //id
	BossId          int     `db:"int(11)"`             //boss id
	VictoryTimes    int     `db:"int(11)"`             //胜利次数
	AddPercent      float64 `db:"float(255,0)"`        //增加的百分比，小数
	AddValue        int64   `db:"bigint(255)"`         //增加的值
	ResourcePercent float64 `db:"float(255,0)"`        //资源系数
}
