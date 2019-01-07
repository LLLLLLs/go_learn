/*
Created on 2018-11-28 13:35:59
author: Auto Generate
*/
package stat

type TextClientHans struct {
	Id    int    `db:"int(255)"`              //id
	Key   string `model:"pk"  db:"char(128)"` //
	Value string `db:"varchar(255)"`          //
	Extra string `db:"varchar(255)"`          //
}
