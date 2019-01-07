/*
Created on 2018-08-16 16:14:42
author: Auto Generate
*/
package stat

type AnabasisChapter struct {
	Id                 int    `model:"pk"  db:"int(11)"` //章id
	BossAttackPower    int64  `db:"bigint(19)"`          //boss攻击力
	RetinueAttackPower int64  `db:"bigint(19)"`          //
	BossAttackMode     int    `db:"int(11)"`             //boss攻击方式(1单体，2横排，3纵排，4全体)
	RetinueAttackMode  int    `db:"int(11)"`             //retinue攻击方式(1单体，2横排，3纵排，4全体)
	BossBlood          int64  `db:"bigint(19)"`          //
	RetinueBlood       int64  `db:"bigint(19)"`          //
	GroupId            string `db:"varchar(64)"`         //
	AdventureTimes     int    `db:"int(255)"`            //新增探索点
}
