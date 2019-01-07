//errors
//created: 2018/7/23
//author: wdj

package errors

var (
	ErrNoTimes             = New("no_times")            //邀约次数不足
	ErrNoSuchSkill         = New("no_such_skill")       //技能不存在
	ErrSkillLocked         = New("skill_locked")        //技能未解锁
	ErrMaxBeautySkillLV    = New("max_beauty_skill_lV") //技能等级达到上限
	ErrInsufficientExp     = New("insufficient_exp")    //经验不足
	ErrBeautyLocked        = New("beauty_locked")       //名媛未解锁
	ErrBeautyExist         = New("beauty_exist")        //名媛存在
	ErrNoVIPLevel          = New("no_vip_level")        //vip等级不足
	ErrNoSuchBeauty        = New("no_such_beauty")      //名媛不存在
	ErrNoSuchOption        = New("no_such_option")      //没有此选项
	ErrNeedNotRecover      = New("need_not_recover")    //不需要恢复
	ErrStudentAccepted     = New("student_accepted")    //学员已接受
	ErrGainAward           = New("gain_award")          //奖励已领取
	ErrMaxBeautyLV         = New("max_beauty_lV")       //最大妃位等级
	ErrBeautyLVNumberLimit = New("num_of_beauty_limit") //同妃位人数限制
	ErrMinBeautyLV         = New("min_beauty_lV")       //最小妃位等级
	ErrMaxChildSeat        = New("max_child_seat")      //最大子嗣席位
	ErrNoSuchSeat          = New("no_such_seat")        //席位不存在
)
