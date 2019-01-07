/*
Author      : lls
Time        : 2018/11/16
Description :
*/

package StudentEnum

type ExamPeriod int16

func (e ExamPeriod) ToInt16() int16 {
	return int16(e)
}

const (
	Middle ExamPeriod = 1 // 期中
	Final  ExamPeriod = 2 // 期末
)
