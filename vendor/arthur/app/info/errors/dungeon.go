/*
Author      : lls
Time        : 2018/09/06
Description :
*/

package errors

var (
	ErrDungeonSingleClosed  = New("dungeon_single_closed")    // 单人副本已关闭
	ErrNoSingleConfig       = New("no_single_config")         // 无单人副本配置
	ErrNoChestConfig        = New("no_chests_config")         // 无宝箱配置
	ErrAllDragonDefeated    = New("all_dragon_defeated")      // 已击败所有巨龙
	ErrInBattling           = New("in_battling")              // 在巨龙战中
	ErrInExploring          = New("in_exploring")             // 在探索中
	ErrBloodNotEnough       = New("blood_not_enough")         // 血量不足
	ErrHeroExhausted        = New("hero_exhausted")           // 英雄出战次数耗尽
	ErrHeroAvailable        = New("hero_available")           // 英雄仍可出战
	ErrHeroNotInExploreList = New("hero_not_in_explore_list") // 英雄不在探索列表中
	ErrHeroStillAlive       = New("hero_still_alive")         // 英雄仍存活
	ErrHeroNotAvailable     = New("hero_not_available")       // 英雄不可出战
	ErrPurchaseTimesLimit   = New("purchase_times_limit")     // 宝箱购买次数达到限制
	ErrHeroRepeat           = New("hero_repeat")              // 英雄重复
	ErrAlreadyExplored      = New("already_explored")         // 该地点已探索

	ErrDungeonGlobalClosed    = New("dungeon_global_closed")     // 全局副本已关闭
	ErrDungeonGlobalHasKilled = New("dungeon_closed")            // 巨龙已被击杀
	ErrDungeonRankNotRegister = New("dungeon_rank_not_register") // 获取副本排行错误：未注册
)
