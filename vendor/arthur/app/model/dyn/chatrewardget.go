/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import (
	"arthur/app/enum/AwardType"
	"arthur/app/infra"
)

type ChatRewardGet struct {
	infra.Model   `model:"-"`
	Id            string         `model:"pk"  db:"char(36)"` //
	RewardId      string         `db:"char(36)"`             //奖励id
	RewardSource  int16          `db:"smallint(11)"`         //奖励来源
	ReceiveRoleId string         `db:"char(36)"`             //获取该奖励的玩家
	RewardType    AwardType.Type `db:"smallint(11)"`         //奖励物品类型
	RewardNo      int16          `db:"smallint(11)"`         //奖励物品编号
	RewardCount   int64          `db:"bigint(36)"`           //奖励物品数量
	Time          int64          `db:"bigint(36)"`           //获取时间
}
