/*
Created on 2018-11-09 10:23:24
author: Auto Generate
*/
package stat

type MarriageProposalLetter struct {
	Id     int16  `model:"pk"  db:"smallint(5)"` //文本id
	IsMale bool   `db:"tinyint(2)"`              //是否为男性
	Text   string `db:"varchar(640)"`            //文本内容
}
