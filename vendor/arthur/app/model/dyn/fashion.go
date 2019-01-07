/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type Fashion struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //时装uuid
	RoleId      string `db:"char(36)"`             //所属玩家id
	No          int16  `db:"smallint(5)"`          //时装编号
	TimeLimit   int64  `db:"bigint(20)"`           //时装过期时间
	Declaration string `db:"char(255)"`            //宣言
}
