/*
Created on 2018-11-06 16:58:42
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type StudentExam struct {
	infra.Model `model:"-"`                        //
	StudentId   string `model:"pk"  db:"char(36)"` //考试学生id
	RoleId      string `db:"char(36)"`             //玩家id
	Stamina     int16  `db:"smallint(5)"`          //体力
	Brain       int16  `db:"smallint(5)"`          //脑力
	Mental      int16  `db:"smallint(5)"`          //心力
	Subject     int16  `db:"smallint(5)"`          //科目
	Round       int16  `db:"smallint(5)"`          //第几题
	ItemUsed    bool   `db:"tinyint(2)"`           //该考试是否使用过道具
	Option11    int16  `db:"smallint(5)"`          //1-1
	Option12    int16  `db:"smallint(5)"`          //1-2
	Option13    int16  `db:"smallint(5)"`          //1-3
	Question1Id int16  `db:"smallint(5)"`          //第一题id
	Option21    int16  `db:"smallint(5)"`          //2-1
	Option22    int16  `db:"smallint(5)"`          //2-2
	Option23    int16  `db:"smallint(5)"`          //2-3
	Question2Id int16  `db:"smallint(5)"`          //第二题id
	Option31    int16  `db:"smallint(5)"`          //3-1
	Option32    int16  `db:"smallint(5)"`          //3-2
	Option33    int16  `db:"smallint(5)"`          //3-3
	Question3Id int16  `db:"smallint(5)"`          //第三题id
}
