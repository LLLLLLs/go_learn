/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type ChatShare struct {
	infra.Model `model:"-"`
	RoleId      string `model:"pk"  db:"char(36)"` //
	ShareTyp    int16  `db:"smallint(11)"`         //分享类型（1英雄，2名媛，3学员）
	ShareTime   int64  `db:"bigint(36)"`           //分享时间
}
