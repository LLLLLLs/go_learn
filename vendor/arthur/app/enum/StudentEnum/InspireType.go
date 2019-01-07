/*
Author      : lls
Time        : 2018/11/15
Description :
*/

package StudentEnum

type InspireType int16

func (i InspireType) ToInt16() int16 {
	return int16(i)
}

const (
	Random InspireType = 1 // 随机鼓励
	Spec   InspireType = 2 // 指定鼓励
)
