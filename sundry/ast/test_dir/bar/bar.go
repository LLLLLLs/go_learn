// @author: lls
// @date: 2021/8/18
// @desc:

package bar

import (
	"fmt"
	"golearn/sundry/ast/values"
)

// test struct
type Bar struct {
	A string // hello string
	B int    // world int
}

func NewService(t testType) Bar {
	call("123456", values.Arg{})
	// 123456:asdfzxcv
	t.methodCall("123456", Arg2{}, Arg{})
	return Bar{}
}

// desc:hello world
// method:comment
func (f Bar) Func1(arg Arg2) {
	fmt.Println(arg)
}

func call(str string, arg interface{}) {

}

type testType struct {
}

func (t testType) methodCall(str string, arg2 interface{}, arg3 interface{}) {

}

type Arg2 struct {
	A   string
	B   int
	C   []float32
	Arg []Arg
}
