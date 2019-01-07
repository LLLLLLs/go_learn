/*
Author  : lls
Time    : 2018/08/15
*/

package errors

var (
	ErrLevyType          = New("levy_type")            // 征收类型错误
	ErrLevyEventNotExist = New("levy_event_not_exist") // 征收事件不存在
	//levy type
	ErrNoLevyType            = New("no_levy_type")            // 征收类型错误
	ErrInsufficientLevyTimes = New("insufficient_levy_times") // 征收次数不足

	ErrCannotOneKeyLevy = New("lv_not_enough_to_one_key_levy") // 等级不足以解锁一键征收
)
