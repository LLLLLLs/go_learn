/*
Created on 2018-09-14 16:04:25
author: Auto Generate
*/
package stat

type ChatReward struct {
	RewardSource int16  `db:"smallint(11)"` //奖励来源
	Name         string `db:"varchar(255)"` //奖励来源
	OneMax       int16  `db:"smallint(11)"` //一个奖励可以被几个玩家领取
	AwardId      string `db:"varchar(255)"` //奖励id
}
