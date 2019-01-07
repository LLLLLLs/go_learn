/*
Author  : lls
Time    : 2018/07/19
*/

package errors

var (
	ErrBattleIllegal   = New("battle_illegal")     // 战斗非法
	ErrBattleEndEarly  = New("battle_end_early")   // 战斗已结束
	ErrBothAlived      = New("both_alived")        // 双方均存活
	ErrWrongSkill      = New("wrong_skill")        // 技能错误
	ErrIllegalSkillCD  = New("illegal_skill_cd")   // 技能CD错误
	ErrOutOfRoundLimit = New("out_of_round_limit") // 超出回合限制
)
