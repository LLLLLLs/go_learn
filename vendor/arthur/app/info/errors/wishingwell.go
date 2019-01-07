/*
Author      : lls
Time        : 2018/08/28
Description :
*/

package errors

// miracle
var (
	ErrNoMiracleConfig = New("no_miracle_config") // 无神迹配置
)

// wish
var (
	ErrNoWishAward      = New("no_wish_award")      // 许愿配置错误
	ErrNoTotalAward     = New("no_total_award")     // 累计奖励配置错误
	ErrAlreadyGetAward  = New("already_get_award")  // 奖励已领取
	ErrNotEnoughDays    = New("not_enough_days")    // 天数不足，领取累计奖励失败
	ErrAlreadyWishToday = New("already_wish_today") // 今日已签到
	ErrWrongWishChoice  = New("wrong_wish_choice")  // 签到选项错误
)

// month card
var (
	ErrWrongMonthCardTye    = New("wrong_month_card_type")  // 月卡类型错误
	ErrNoMonthCard          = New("no_month_card")          // 无月卡
	ErrMonthCardOutDated    = New("month_card_out_dated")   // 月卡已过期
	ErrAlreadyReceivedToday = New("already_received_today") // 今日已领取
)
