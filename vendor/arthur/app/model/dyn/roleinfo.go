/*
Created on 2018-11-06 16:58:42
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type RoleInfo struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	UserId      string `db:"char(36)"`             //用户id
	ShowId      int    `db:"int(10)"`              //客户端内显示id
	Nickname    string `db:"char(10)"`             //昵称
}
