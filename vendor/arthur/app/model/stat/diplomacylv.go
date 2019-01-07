/*
Created on 2018-11-08 15:26:08
author: Auto Generate
*/
package stat

type DiplomacyLv struct {
	Lv          int     `model:"pk"  db:"int(255)"` //
	Require     int     `db:"int(11)"`              //外交点需求
	Rate        float64 `db:"float(11,2)"`          //外交等级国力加成
	Attr        int64   `db:"bigint(255)"`          //外交等级属性
	AwardWeight float64 `db:"float(11,2)"`          //外交点奖励加成
}
