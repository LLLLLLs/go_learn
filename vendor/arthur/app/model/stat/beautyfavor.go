/*
Created on 2019-01-04 14:29:17
author: Auto Generate
*/
package stat

type BeautyFavor struct {
	Id    int    `model:"pk"  db:"int(11)"` //物品id
	Name  string `db:"varchar(255)"`        //物品名称
	PosId int    `db:"int(11)"`             //物品摆放位置编号，具体位置参照预制体
}
