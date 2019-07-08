// Time        : 2019/07/08
// Description : 空结构体

package _struct

import "fmt"

type empty struct{}

func (e empty) hello() {
	fmt.Println("hello")
}

func (e empty) world() {
	fmt.Println("world")
}

// 可以用作map的value使用
var hasKey = make(map[string]struct{})

// channel传递不需要携带数据的时候
var c = make(chan struct{})

func init() {
	_ = hasKey
	_ = c
}
