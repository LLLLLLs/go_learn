/*
Created on 2018-12-26 10:01:57
author: Auto Generate
*/
package stat

type VisitEntrustEvents struct {
	Id        int    `model:"pk"  db:"int(11)"` //事件编号
	Name      string `db:"varchar(255)"`        //事件名字
	Attr      int    `db:"int(11)"`             //属性类型要求
	Intro     string `db:"varchar(255)"`        //事件简介
	CountryId int    `model:"pk"  db:"int(11)"` //国家id， 0=通用型
}
