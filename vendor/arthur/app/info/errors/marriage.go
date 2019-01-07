/*
Author      : lls
Time        : 2018/10/10
Description :
*/

package errors

var (
	// 联姻状态
	ErrStudentAlreadyMarried = New("student_already_married") // 学员已联姻
	ErrStudentInSeeking      = New("student_in_seeking")      // 学员正在寻缘
	ErrStudentInProposing    = New("student_in_proposing")    // 学员正在求婚
	ErrWrongProposingTarget  = New("wrong_proposing_target")  // 求婚目标错误
	// 学员状态
	ErrBothTheSameSex       = New("both_the_same_sex")        // 双方同性
	ErrMedalNotMatch        = New("medal_not_match")          // 等级不符
	ErrBelongsToTheSameRole = New("belongs_to_the_same_role") // 同一角色
	// 联姻类型
	ErrWrongMarriageType = New("wrong_marriage_type") // 联姻类型错误
	// 时间状态
	ErrSeekingOutDated  = New("seeking_out_dated")  // 寻缘超时
	ErrProposalOutDated = New("proposal_out_dated") // 求婚超时
	// 数据库记录
	ErrStudentNotInProposing = New("student_not_in_proposing") // 学员不在求婚中
	ErrStudentNotInSeeking   = New("student_not_in_seeking")   // 学员不在寻缘中
	ErrNoProposingRecord     = New("no_proposing_record")      // 没有该条求婚记录
	// 配置
	ErrNoProposalLetter = New("no_proposal_letter") // 无求婚信配置
	// 求婚
	ErrCanNotProposeSelf = New("can_not_propose_self") // 无法向自己求婚
)
