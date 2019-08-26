// Time        : 2019/08/22
// Description :

package function

import "fmt"

func NewFunc() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}

type ContentFunc struct {
	f func() int
}

func (cf ContentFunc) Do() {
	fmt.Println(cf.f())
}
