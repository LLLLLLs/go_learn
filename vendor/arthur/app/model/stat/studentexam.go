/*
Created on 2018-11-15 15:50:47
author: Auto Generate
*/
package stat

type StudentExam struct {
	QuestionId   int16 `model:"pk"  db:"smallint(5)"` //题目id
	Subject      int16 `db:"smallint(5)"`             //科目编号
	QuestionList []int `db:"varchar(10)"`             //题目类型(区分正负)
}
