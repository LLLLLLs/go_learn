// Time        : 2019/06/26
// Description :

package _defer

import (
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
