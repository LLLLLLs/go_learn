/*
Created on 2018-11-16 10:57:08
author: Auto Generate
*/
package dyn

import (
	"arthur/app/enum/StudentEnum"
	"arthur/app/infra"
)

type StudentExamResult struct {
	infra.Model   `model:"-"`                                        //
	Id            string                 `model:"pk"  db:"char(36)"` //
	StudentId     string                 `db:"char(36)"`             //学员ID
	Subject       int16                  `db:"smallint(5)"`          //考试科目
	Period        StudentEnum.ExamPeriod `db:"smallint(5)"`          //时期(期中、期末) 1=期中 2=期末
	Result        bool                   `db:"tinyint(1)"`           //考试结果
	Time          int64                  `db:"bigint(20)"`           //考试结束时间
	WiseDelta     int                    `db:"int(10)"`              //智慧属性变化量
	DiligentDelta int                    `db:"int(10)"`              //勤勉属性变化量
	LoyaltyDelta  int                    `db:"int(10)"`              //忠诚属性变化量
	HeroicDelta   int                    `db:"int(10)"`              //英勇属性变化量
}
