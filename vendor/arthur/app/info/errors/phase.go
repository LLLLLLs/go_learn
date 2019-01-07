/*
Author: Minsi Ruan
Data: 2018/7/16 16:38
*/

package errors

var (
	ErrLackSolidersOrAttack = New("err_lack_soliders_or_attack") // 缺少士兵或攻击力
	ErrNoPhase              = New("err_no_phase")                // 关卡信息错误
)
