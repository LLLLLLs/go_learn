/*
Created on 2018-12-10 11:55:51
author: Auto Generate
*/
package stat

type BeautyLvChild struct {
	SeatNum  int `model:"pk"  db:"int(11)"` //席位数量
	Required int `db:"int(11)"`             //需要的钻石
}
