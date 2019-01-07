/*
Created on 2018-08-17 16:24:57
author: Auto Generate
*/
package stat

type RoleVip struct {
	Level        int16  `model:"pk"  db:"smallint(10)"` //玩家vip等级
	Pregnant     int    `db:"int(10)"`                  //子嗣诞生概率加成
	Vitality     int    `db:"int(10)"`                  //活力上限
	Vigor        int    `db:"int(10)"`                  //精力上限
	Act          int    `db:"int(10)"`                  //行动力上限
	Miracle      int    `db:"int(10)"`                  //神迹触发次数
	GoldenChest  int16  `db:"smallint(5)"`              //黄金宝箱限购
	SilverChest  int16  `db:"smallint(5)"`              //白银宝箱限购
	Cost         int    `db:"int(10)"`                  //充值需求
	AwardId      string `db:"char(36)"`                 //vip福利
	InspireTimes int16  `db:"smallint(5)"`              //修道院鼓励次数
}
