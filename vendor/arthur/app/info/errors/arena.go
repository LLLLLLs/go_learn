/*
Author  : lls
Time    : 2018/07/20
*/

package errors

var (
	ErrNotEnoughHeroesToUnlock1v1 = New("not_enough_heroes_to_unlock_1v1") // 英雄不足以解锁1v1

	ErrNoSoloSkillConf   = New("no_battle_skill_conf") // 无技能配置
	ErrHeroCannotFight   = New("hero_cannot_fight")    // 英雄无法出战
	ErrHeroStillCanFight = New("hero_still_can_fight") // 英雄仍可出战

	ErrNoUnAcceptedSolo = New("no_unaccepted_solo") // 无未接受战斗
	ErrAlreadyAccept    = New("already_accept")     // 战斗已接受
	ErrRecordNotAccept  = New("record_not_accept")  // 战斗还未接受
	ErrNoArenaNPC       = New("no_arena_npc")       // 无竞技场NPC配置
	ErrCannotFightSelf  = New("cannot_fight_self")  // 不能挑战自己

	ErrNoCurrentRecord = New("no_current_record") // 当前无记录
	ErrNoThisFlow      = New("no_this_flow")      // 无该条流水
	ErrNoThisDeffence  = New("no_this_deffence")  // 无该防守记录
	ErrWrongTarget     = New("wrong_target")      // 目标错误

	ErrAlreadyInSoloBattling = New("already_in_solo_battling") // 已在战斗中

	ErrStillHaveRandomTimes = New("still_have_free_random_times") // 免费次数未耗尽
	ErrRandomTimesExhausted = New("random_times_exhausted")       // 免费次数已耗尽

	// store
	ErrNoItemInArenaStore = New("no_item_in_arena_store") // 竞技场商店中无该道具
	ErrNoArenaCoin        = New("no_arena_coin")          // 竞技币不足
	ErrArenaItemSoldOut   = New("arena_item_sold_out")    // 竞技场商品售罄
)
