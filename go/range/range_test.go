// Time        : 2019/11/11
// Description :

package _range

import (
	"fmt"
	"testing"
)

// 闭包直接使用range的返回值 -- 实际上是指针
func TestRange(t *testing.T) {
	var list = []int{1, 2, 3, 4, 5}
	var fs = make([]func() (int, int), len(list))
	for i, v := range list {
		fs[i] = func() (int, int) {
			return i, v
		}
	}
	for i := range fs {
		fmt.Println(fs[i]())
	}
}

// 多一层函数 -- 多一个值传递的操作
func TestRange1(t *testing.T) {
	var list = []int{1, 2, 3, 4, 5}
	var fs = make([]func() (int, int), len(list))
	for i, v := range list {
		fs[i] = func(i, v int) func() (int, int) {
			return func() (int, int) {
				return i, v
			}
		}(i, v)
	}
	for i := range fs {
		fmt.Println(fs[i]())
	}
}

// 手动值拷贝
func TestRange2(t *testing.T) {
	var list = []int{1, 2, 3, 4, 5}
	var fs = make([]func() (int, int), len(list))
	for i, v := range list {
		index, value := i, v
		fs[i] = func() (int, int) {
			return index, value
		}
	}
	for i := range fs {
		fmt.Println(fs[i]())
	}
}
