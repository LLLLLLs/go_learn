/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type MarriageTemp struct {
	infra.Model   `model:"-"`                        //
	StudentId     string `model:"pk"  db:"char(36)"` //待处理学员id
	RoleId        string `db:"char(36)"`             //待处理玩家id
	TargetRole    string `db:"char(36)"`             //联姻目标玩家id
	TargetStudent string `db:"char(36)"`             //联姻目标学员id
	Time          int64  `db:"bigint(20)"`           //联姻成功时间
}
