/*
Author: Minsi Ruan
Data: 2018/11/13 10:22
*/

package errors

var (
	ErrNotHeroNPC  = New("err_not_hero_npc")       // 没有英雄NPC
	ErrUnAchieve   = New("err_un_get_achievement") //成就未达成
	ErrProposalNum = New("err_proposal_num")       // 议案数量错误
	ErrLackOfEXP   = New("err_lack_of_exp")        //缺少外交经验
	ErrAttr        = New("err_attr")               // 属性错误
	ErrLackOfAct   = New("err_lack_of_act")        // 缺少议案处理次数
	ErrMaxLv       = New("err_max_lv")             //已达到最高外交等级
	ErrOverUse     = New("err_over_use")           //道具使用过度
)
