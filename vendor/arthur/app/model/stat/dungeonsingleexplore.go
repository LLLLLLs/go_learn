/*
Created on 2018-09-11 10:45:30
author: Auto Generate
*/
package stat

type DungeonSingleExplore struct {
	No               int16   `model:"pk"` //探索地点编号
	BloodConsumeCoef float64 //探索消耗系数
	ScoreCoef        float64 //探索积分系数
	Award            []int   //宝箱获取概率
}
