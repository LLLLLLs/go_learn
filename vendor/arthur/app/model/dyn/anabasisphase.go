/*
Created on 2018-11-06 16:58:40
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type AnabasisPhase struct {
	infra.Model   `model:"-"`
	RoleId        string `model:"pk"  db:"varchar(36)"` //角色id
	Soldiers      int64  `db:"bigint(19)"`              //当前关卡剩余敌军数
	RefreshTime   int64  `db:"bigint(20)"`              //关卡刷新时间
	PlayBoss      int16  `db:"smallint(5)"`             //是否播放过boss动画
	PlaySection   int16  `db:"smallint(5)"`             //是否播放过节动画
	PlayChapter   int16  `db:"smallint(5)"`             //是否播放过章动画
	PhaseStat     int    `db:"int(11)"`                 //关卡状态（1普通关，2boss关）
	BossBlood     int64  `db:"bigint(19)"`              //大boss剩余血量
	RetinueBloodA int64  `db:"bigint(19)"`              //boss随从a剩余血量
	RetinueBloodB int64  `db:"bigint(19)"`              //boss随从b剩余血量
	RetinueBloodC int64  `db:"bigint(19)"`              //boss随从c剩余血量
}
