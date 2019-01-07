/*
Created on 2018-11-09 11:58:24
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type DiplomacyProposal struct {
	infra.Model `model:"-"` //
	Id          string      `model:"pk"  db:"char(36)"` //
	RoleId      string      `db:"char(36)"`             //
	ProposalNo  int         `db:"int(11)"`              //议案编号
	ProposalNpc int         `db:"int(11)"`              //提案NPC
	CountryNo   int         `db:"int(11)"`              //议案国家
	No          int         `db:"int(11)"`              //议案总数（1、2、3）
	Position    int         `db:"int(11)"`              //议案显示位置（1-8）
}
