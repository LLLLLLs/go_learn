/*
Created on 2018-09-20 14:29:55
author: Auto Generate
*/
package stat

type AdventureCity struct {
	City            int     `model:"pk"  db:"int(11)"` //城堡序号
	PositionList    []int   `db:"varchar(255)"`        //城堡地点id列表
	BossAward       [][]int `db:"varchar(255)"`        //boss奖励， 格式[[道具id， 数量]]
	BoxA            [][]int `db:"varchar(255)"`        //1号宝箱奖励， 格式[[道具id， 数量]]
	BoxB            [][]int `db:"varchar(255)"`        //2号宝箱奖励， 格式[[道具id， 数量]]
	BossBlood       int64   `db:"bigint(20)"`          //boss血量
	BossAttack      int64   `db:"bigint(20)"`          //boss攻击力
	ChapterRequired int     `db:"int(6)"`              //需要的解锁章节
}
