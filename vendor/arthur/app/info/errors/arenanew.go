/*
@Time : 2018/10/19 17:24
@Author : linfeng
@File : arenanew
@Desc:
*/

package errors

import "arthur/utils/errors"

var (
	ErrorNotAllowArena         = errors.New("error_not_allow_arena")        //不满足进入竞技场条件
	ErrUnableFindFight         = errors.New("unable_find_fight")            //无法查找战斗
	ErrUnableReceiveFight      = errors.New("unable_receive_fight")         //无法接受战斗
	ErrDeltaUseEnvoyNum        = errors.New("err_delta_use_envoy_num")      //英雄数和出使令使用比例出错
	ErrUnableNewBuff           = errors.New("unable_new_buff")              //生成buff时出错
	ErrNotEnoughMoneyPayBuff   = errors.New("not_enough_money_pay_buff")    //没有足够的宝珠或元宝支付buff
	ErrPayBuffNotMatchCostType = errors.New("not_match_cost_type")          //支付buff时没有匹配到支付方式
	ErrHadBuff                 = errors.New("had_buff")                     //重复购买某个位置上的BUFF
	ErrFighting                = errors.New("err_fighting")                 //正在战斗中，无法开启新战斗
	ErrorWrongStatusFight      = errors.New("error_wrong_status_fight")     //战斗状态错误
	ErrorStartFight            = errors.New("error_start_fight")            //尚未开始战斗
	ErrorRoundWin              = errors.New("error_round_win")              //本轮胜负状态错误
	ErrorNotEnoughItemToUse    = errors.New("error_not_enough_item_to_use") //没有足够的道具可以消耗
	ErrorFightYouself          = errors.New("error_fight_youself")          //不能挑战或复仇自己
	ErrorHadChallenge          = errors.New("had_challenge")                //已挑战或已复仇
	ErrorZeroChallengeHeros    = errors.New("error_zero_challenge_heros")   //挑战或复仇时无可出战的英雄
	ErrorHasUnfinishFighting   = errors.New("error_has_unfinish_fighting")  //还有尚未完成的战斗
	ErrorAlloweFixPlace        = errors.New("error_allowe_fix_place")       //无法修改阵容
	ErrorHeroUnableFight       = errors.New("error_hero_unable_fight")      //英雄无法战斗
	ErrorHadFixPlace           = errors.New("error_had_fix_place")          //已经设置过阵容了
	ErrorHadFirstBuff          = errors.New("error_had_first_buff")         //首轮已购买BUFF
	ErrorAttackingNotValid     = errors.New("error_attacking_not_valid")    //参战英雄都是占位英雄，用于挑战赛
	ErrorRepeatHerosNo         = errors.New("error_repeat_heros_no")        //英雄编号重复
	ErrorReceive               = errors.New("error_receive")                //接受战斗时出错
	ErrorEnvoy                 = errors.New("error_envoy")                  //出使错误
	ErrorReloadNotEnoughGold   = errors.New("error_reload_not_enough_gold") //替换对手时没有足够的元宝
)
