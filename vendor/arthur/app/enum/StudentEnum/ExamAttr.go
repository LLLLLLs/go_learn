/*
Author      : lls
Time        : 2018/11/15
Description :
*/

package StudentEnum

type ExamAttr int16

func (e ExamAttr) ToInt16() int16 {
	return int16(e)
}

const (
	Brain   ExamAttr = 1 // 脑力
	Stamina ExamAttr = 2 // 体力
	Mental  ExamAttr = 3 // 心力
)

var AttrList = []ExamAttr{Brain, Stamina, Mental}
