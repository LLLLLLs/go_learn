/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type Beauty struct {
	infra.Model     `model:"-"` //
	BeautyId        string      `model:"pk"  db:"char(36)"` //佳人id
	RoleId          string      `db:"varchar(36)"`          //rid
	BeautyType      int         `db:"int(11)"`              //佳人类型
	Amity           int         `db:"int(11)"`              //友好度
	Charm           int         `db:"int(11)"`              //魅力值
	Exp             int64       `db:"bigint(11)"`           //佳人经验
	Favor           int         `db:"int(11)"`              //好感度
	IsUnlock        bool        `db:"tinyint(4)"`           //is_unlock
	HintUnlock      bool        `db:"tinyint(4)"`           //是否提示解锁
	SkillA          int16       `db:"smallint(6)"`          //技能1的等级
	SkillB          int16       `db:"smallint(6)"`          //技能2的等级
	SkillC          int16       `db:"smallint(6)"`          //
	SkillD          int16       `db:"smallint(6)"`          //
	SkillE          int16       `db:"smallint(6)"`          //
	SkillF          int16       `db:"smallint(6)"`          //
	EventId         int         `db:"int(11)"`              //邀约剧情id
	OfflineChild    string      `db:"varchar(255)"`         //掉线的学员id
	NextFavorChoice int64       `db:"bigint(255)"`          //下次选择的喜爱物品的随机种子
	Lv              int         `db:"int(255)"`             //妃位等级
}
