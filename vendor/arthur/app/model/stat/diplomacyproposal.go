/*
Created on 2018-11-08 15:26:47
author: Auto Generate
*/
package stat

type DiplomacyProposal struct {
	No        int    `model:"pk"  db:"int(11)"` //
	CountryId int    `db:"int(6)"`              //国家id
	Type      int16  `db:"smallint(6)"`         //主数
	Second    int16  `db:"smallint(6)"`         //辅数
	Desc      string `db:"text"`                //议案描述
	PowerA    int    `db:"int(11)"`             //教会势力
	PowerB    int    `db:"int(11)"`             //军队势力
	PowerC    int    `db:"int(11)"`             //商会势力
	PowerD    int    `db:"int(11)"`             //平民势力
}
