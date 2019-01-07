/*
Author : Haoyuan Liu
Time   : 2018/7/9
*/
package errors

var (
	ErrHeroMaxPeerage = New("max_peerage")   //英雄已升至最大爵位
	ErrLvInsuff       = New("lv_insuff")     //等级不满足升爵条件
	ErrNoAptiLv       = New("no_apti_lv")    //无天赋配置
	ErrTalentLimit    = New("talent_limit")  //超过天赋上限
	ErrNoAptiExp      = New("no_apti_exp")   //天赋经验不足
	ErrNoFaith        = New("no_faith")      //英雄无信仰
	ErrNoFaithSkill    = New("NoFaithSkill")
	ErrFaithSkillMaxLv = New("FaithSkillMaxLv")
	ErrWrongItem      = New("wrong_item")    //错误的道具
	ErrHeroLvLimit    = New("hero_lv_limit") //英雄等级已升至最大
	ErrHeroExist      = New("hero_exist")    //英雄已存在
)
