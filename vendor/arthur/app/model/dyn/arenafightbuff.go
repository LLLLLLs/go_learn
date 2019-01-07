/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ArenaFightBuff struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"char(36)"` //
	FightId     string `db:"char(36)"`             //本场战斗id
	RoundId     string `db:"char(36)"`             //本轮战斗id
	RoleId      string `db:"char(36)"`             //角色ID
	Status      int16  `db:"smallint(6)"`          //本场状态：0战斗未结束，1战斗已结束
	RoundWin    int16  `db:"smallint(6)"`          //本轮状态：0尚未确定；1胜利；2失败
	PositionNo  int16  `db:"smallint(6)"`          //buff牌编号
	BuffType    int16  `db:"smallint(6)"`          //buff类型：0为攻击，1为血量
	BuffValue   int16  `db:"mediumint(9)"`         //buff效果对应的值
	BuyTime     int64  `db:"bigint(20)"`           //购买buff时间
	CostType    int16  `db:"smallint(6)"`          //购买buff花费类型：0为钻石，1为宝珠
	Amount      int16  `db:"mediumint(9)"`         //花费
}
