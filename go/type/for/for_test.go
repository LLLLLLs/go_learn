//@author: lls
//@time: 2021/02/07
//@desc:

package _for

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	forRange(3, 4)
	forRange(4, 4)
	forRange(4, 3)
}

func forRange(a, b int) {
	for i := a + 1; i <= b; i++ {
		fmt.Println(i)
	}
}
