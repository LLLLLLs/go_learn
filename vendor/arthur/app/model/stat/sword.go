/*
Created on 2018-10-19 17:42:06
author: Auto Generate
*/
package stat

type Sword struct {
	No             int        `model:"pk"  db:"int(11)"` //王者之剑编号
	Chapter        int        `db:"int(11)"`             //解锁关卡数
	Country        []int      `db:"varchar(255)"`        //需要加成的国家
	Resource       [3]float64 `db:"varchar(255)"`        //资源加成百分比,[金币，粮食， 士兵]
	LossSoldier    float64    `db:"float(255,2)"`        //减少兵损
	AnabasisAttack float64    `db:"float(255,2)"`        //远征boss战上阵英雄攻击加成百分比
	AnabasisBlood  float64    `db:"float(255,2)"`        //远征boss战上阵英雄血量加成百分比
	ArenaAttack    float64    `db:"float(255,2)"`        //竞技场上阵英雄攻击加成加成百分比
	ArenaBlood     float64    `db:"float(255,2)"`        //竞技场上阵英雄血量加成加成百分比
	ArenaScore     float64    `db:"float(255,2)"`        //竞技场中获得的竞技分数加成百分比
	ArenaCoin      float64    `db:"float(255,2)"`        //竞技场中获得的竞技币加成百分比
	Diplomacy      float64    `db:"float(255,2)"`        //圆桌厅加成
	Student        []float64  `db:"varchar(255)"`        //学院加成列表
	BeautyLv       []float64  `db:"varchar(255)"`        //名媛妃位(10级)加成列表
	DragonAttack   float64    `db:"float(255,2)"`        //副本上阵英雄攻击
	DragonBlood    float64    `db:"float(255,2)"`        //副本上阵英雄血量
	DragonScore    float64    `db:"float(255,2)"`        //副本分数
	DragonCoin     float64    `db:"float(255,2)"`        //副本龙币
}
