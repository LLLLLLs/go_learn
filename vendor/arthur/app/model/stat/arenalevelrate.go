/*
Created on 2018-10-19 16:32:49
author: Auto Generate
*/
package stat

type ArenaLevelRate struct {
	No              int16 `model:"pk"` //
	LowerTime       int16 //翻牌最低次数
	HeightTime      int16 //翻牌最高次数
	LowerLevelNo    int16 //较低档次编号
	HeightLevelNo   int16 //较高档次编号
	LowerLevelRate  int16 //较低档次概率
	HeightLevelRate int16 //较高档次概率
	AttackRate      int16 //攻击buff概率
	BloodRate       int16 //血量buff概率
}
