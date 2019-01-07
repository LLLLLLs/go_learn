/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type KeyValue struct {
	infra.Model `model:"-"`
	DynKey      string `model:"pk"  db:"varchar(255)"` //
	Value       string `db:"varchar(255)"`             //
	ValueType   string `db:"varchar(255)"`             //
	Desc        string `db:"varchar(255)"`             //
}
