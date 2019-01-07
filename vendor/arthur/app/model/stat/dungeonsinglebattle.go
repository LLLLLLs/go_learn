/*
Created on 2018-09-11 10:45:26
author: Auto Generate
*/
package stat

type DungeonSingleBattle struct {
	Level       int16   `model:"pk"` //关卡数
	Blood       int64   //boss血量
	ExploreCoef float64 //探索关系系数
	Attack      int64   //boss攻击力
	TotalScore  int16   //关卡总积分
	BattleScore int16   //战胜积分
	Profession  int16   //boss职业
}
