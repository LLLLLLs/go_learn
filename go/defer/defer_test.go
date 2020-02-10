// Time        : 2019/06/26
// Description :

package _defer

import (
	"errors"
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	a := 1
	defer fmt.Println(a)

	defer d(func() {
		fmt.Println(a)
	})

	defer func() {
		fmt.Println(a)
		a = 3
	}()
	a = 2
}

func d(f func()) {
	f()
}

func TestRecover(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	panic(111)
}

// defer闭包
func TestDefer_Clause(t *testing.T) {
	err := deferClause()
	fmt.Println("test:", err)
}

// defer直接调用
func TestDefer_Exec(t *testing.T) {
	err := deferExec()
	fmt.Println("test:", err)
}

func deferClause() (err error) {
	defer func() { handle(err) }()
	err = errors.New("234")
	return
}

func deferExec() (err error) {
	defer handle(err)
	err = errors.New("234")
	return
}

func handle(err error) {
	fmt.Println("defer:", err)
}
