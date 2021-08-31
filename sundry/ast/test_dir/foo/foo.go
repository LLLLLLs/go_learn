// @author: lls
// @date: 2021/8/18
// @desc:

package foo

import (
	"fmt"
	"golearn/sundry/ast/values"
)

// test struct
type Foo struct {
	A string // hello string
	B int    // world int
}

// comment
func (f Foo) Func1(arg values.Arg) {
	fmt.Println(arg)
}
