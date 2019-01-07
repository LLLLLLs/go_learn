/*
Created on 2018-12-04 15:43:25
author: Auto Generate
*/
package stat

type AnabasisPhase struct {
	Chapter    int      `db:"int(11)"`      //章数
	Phase      int      `db:"int(11)"`      //关数
	Name       string   `db:"varchar(255)"` //名称
	Heroic     int64    `db:"bigint(19)"`   //标准武力值
	Soldiers   int64    `db:"bigint(19)"`   //敌军数
	GroupId    string   `db:"varchar(255)"` //奖励组
	EnemyImgId int      `db:"int(11)"`      //
	Award      [][2]int //奖励组
}
