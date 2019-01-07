/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import (
	"arthur/app/enum/Marriage"
	"arthur/app/infra"
)

type MarriageProposal struct {
	infra.Model `model:"-"`                                  //
	StudentId   string           `model:"pk"  db:"char(36)"` //学员id
	RoleId      string           `db:"char(36)"`             //玩家id
	TargetId    string           `db:"char(36)"`             //目标玩家id
	ConsumeType Marriage.Consume `db:"smallint(5)"`          //消耗类型（道具 or 钻石）
	Time        int64            `db:"bigint(20)"`           //求婚时间
	Stat        bool             `db:"tinyint(1)"`           //求婚状态：false=拒绝，true=无
	TextId      int16            `db:"smallint(5)"`          //求婚信文本编号
	IsRead      bool             `db:"tinyint(1)"`           //是否读过
}
