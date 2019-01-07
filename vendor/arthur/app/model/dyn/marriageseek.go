/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import (
	"arthur/app/enum/Marriage"
	"arthur/app/enum/StudentEnum"
	"arthur/app/infra"
)

type MarriageSeek struct {
	infra.Model `model:"-"`                                   //
	StudentId   string            `model:"pk"  db:"char(36)"` //学员ID
	RoleId      string            `db:"char(36)"`             //玩家ID
	Type        StudentEnum.Medal `db:"smallint(5)"`          //学员等级
	ConsumeType Marriage.Consume  `db:"smallint(5)"`          //道具消耗类型（钻石 or 道具）
	Sex         bool              `db:"tinyint(2)"`           //学员性别
	Time        int64             `db:"bigint(20)"`           //寻缘开始时间
}
