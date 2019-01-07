/*
Author: Minsi Ruan
Data: 2018/8/13 9:14
*/

package errors

var (
	ErrActEnough         = New("act_enough")           //行动力非空
	ErrActNoEnough       = New("no_enough")            //行动力不足
	ErrNPCNotMeet        = New("not_meet_npc")         //未遇见npc
	ErrNoSuchHero        = New("no_such_hero")         //英雄不存在
	ErrNoSuchCountry     = New("no_such_country")      //国家不存在
	ErrNoSuchNPC         = New("no_such_npc")          //npc不存在
	ErrNoSuchEvent       = New("no_such_event")        //事件不存在
	ErrUnknownStage      = New("unknown_stage")        //未知阶段
	ErrMaxEntrustEventLV = New("max_entrust_event_lv") //最大委派事件等级
	ErrAttrInsufficient  = New("insufficient_attr")    //属性不足
	ErrNeedNotRecovery   = New("need_not_recovery")    //不需要恢复
	ErrNoDispatchTimes   = New("no_dispatch_times")    //派遣次数用完
	ErrMoreResource      = New("more_resource")        //资源不足
	ErrAlreadyRecovery   = New("already_recovery")     //已经恢复
	ErrInDispatch        = New("in_dispatch")          //英雄在派遣中
	ErrRequireChapter    = New("pass_more_chapter")    //解锁更多章节
	ErrNotTheCountry     = New("err_not_the_country")  //未在国家中
)
