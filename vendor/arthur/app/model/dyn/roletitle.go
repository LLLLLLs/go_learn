/*
Created on 2018-11-06 16:58:42
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type RoleTitle struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	RoleId      string `db:"char(36)"`             //
	No          int16  `db:"smallint(5)"`          //称号编号
}
