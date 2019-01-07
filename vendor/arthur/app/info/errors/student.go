/*
Author  : lls
Time    : 2018/07/16
*/

package errors

var (
	// student
	ErrStudentNotExist = New("student_not_exist")   // 学员不存在
	ErrCannotGraduate  = New("cannot_graduate")     // 学员无法毕业
	ErrNoMother        = New("can_not_find_mother") // 找不到名媛

	// config
	ErrNoVipConf               = New("no_vip_conf")                // 无vip配置
	ErrNoCloisterConfig        = New("no_cloister_config")         // 无修道院配置
	ErrNoStudentTalentConfig   = New("no_student_talent_config")   // 学员天赋配置错误
	ErrNoStudentMedalConfig    = New("no_student_medal_config")    // 学员勋章配置错误
	ErrNoCloisterSubjectConfig = New("no_cloister_subject_config") // 修道院科目配置错误
	ErrNoStudentAvatarConfig   = New("no_student_avatar_config")   // 学员形象配置错误

	// cloister
	ErrCloisterExpNotEnough      = New("cloister_exp_not_enough")       // 修道院经验不足
	ErrCloisterLvMax             = New("cloister_lv_max")               // 修道院已满级
	ErrCannotUnlockOneKeyInspire = New("cannot_unlock_one_key_inspire") // 无法解锁一键鼓励
	ErrNoInspireTimes            = New("no_inspire_times")              // 无鼓励次数

	// exam
	ErrIllegalExamChoice      = New("illegal_exam_choice")        // 考试选择非法
	ErrExamNotDefeat          = New("exam_not_defeat")            // 考试未失败
	ErrStudentExaming         = New("student_examing ")           // 学员考试中
	ErrStudentNotReadyForExam = New("student_not_ready_for_exam") // 学员无法考试
	ErrCloisterNotHasSubject  = New("cloister_not_has_subject")   // 修道院未解锁该科目
	ErrStudentNotInAnExam     = New("student_not_in_an_exam")     // 学员不在考试中
	ErrExamFailed             = New("exam_failed")                // 考试失败
	ErrExamItemUsed           = New("exam_item_used")             // 考校宝典已使用

	// inspire
	ErrCannotInspire       = New("cannot_inspire")         // 无法鼓励
	ErrStudentNotEscape    = New("student_not_es_cape")    // 学员未逃学
	ErrNoCanInspireStudent = New("no_can_inspire_student") // 无可鼓励学员
	ErrStillCanInspire     = New("still_can_inspire")      // 仍有鼓励次数
)
