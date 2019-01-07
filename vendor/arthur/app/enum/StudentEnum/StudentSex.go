/*
Author      : lls
Time        : 2018/10/11
Description : 
*/

package StudentEnum

type Sex bool

func (sex Sex) Bool() bool {
	return bool(sex)
}

var (
	Male   Sex = true
	Female Sex = false
)
