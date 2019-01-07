/*
Created on 2018-12-10 14:26:58
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type BeautyChild struct {
	infra.Model `model:"-"` //
	Id          string      `model:"pk"  db:"char(36)"` //
	BeautyId    string      `db:"char(36)"`             //佳人id
	StudentId   string      `db:"char(36)"`             //学员id
	Index       int         `db:"int(11)"`              //席位id，从1开始计
}
