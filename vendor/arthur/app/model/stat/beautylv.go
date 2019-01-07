/*
Created on 2018-12-10 11:55:48
author: Auto Generate
*/
package stat

type BeautyLv struct {
	BeautyLv      int     `model:"pk"  db:"int(11)"` //妃位等级
	ChipRequired  int     `db:"int(11)"`             //需要的碎片
	Name          string  `db:"varchar(255)"`        //等级名称
	ChildPercent  float64 `db:"float(255,2)"`        //子嗣转化百分比
	NumLimit      int     `db:"int(11)"`             //名媛数量限制
	DowngradeRate float64 `db:"float"`               //降级到下级返还的百分比
}
