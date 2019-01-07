/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import (
	"arthur/app/enum/StudentEnum"
	"arthur/app/infra"
)

type MarriageHall struct {
	infra.Model `model:"-"`
	Id          string            `model:"pk"  db:"char(36)"` //
	RoleId      string            `db:"char(36)"`             //角色id
	Type        StudentEnum.Medal `db:"smallint(5)"`          //学员爵位等级
	Sex         bool              `db:"tinyint(5)"`           //学员性别
	Record1     string            `db:"char(36)"`             //第一条寻缘记录
	Record2     string            `db:"char(36)"`             //第二条寻缘记录
	Record3     string            `db:"char(36)"`             //第三条寻缘记录
	Record4     string            `db:"char(36)"`             //第四条寻缘记录
	RefreshTime int64             `db:"bigint(20)"`           //刷新时间
}
