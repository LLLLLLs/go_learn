/*
Created on 2018/7/17 9:23

author: ChenJinLong

Content:
*/
package errors

var (
	ErrHeroHaveFought       = New("err_hero_have_fought")        // 英雄已经战斗过
	ErrNoBossPhase          = New("err_no_boss_phase")           // 本章节没有Boss
	ErrNoBattleField        = New("err_no_battle_field")         // 没有多余的战斗坑位
	ErrCannotRestoreFight   = New("err_cannot_restore_fight")    // 无法重新开战
	ErrNoProfConf           = New("err_no_prof_conf")            // 英雄天赋配置错误
	ErrNoHero               = New("err_no_hero")                 // 没有英雄
	ErrHeroForRoleIncorrect = New("err_hero_for_role_incorrect") // 英雄与角色不符合
)
