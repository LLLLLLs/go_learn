/*
Author: Minsi Ruan
Data: 2018/9/17 14:45
*/

package errors

var (
	ErrAlreadyReceive = New("err_already_received")  // 奖励已领取
	ErrLackOfRequire  = New("err_lack_of_require!")  // 领取条件不足
	ErrLackOfActivity = New("err_lack_of_activity ") // 活跃度不足
)
