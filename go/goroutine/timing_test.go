// Time        : 2019/12/02
// Description :

package goroutine

import (
	"fmt"
	"testing"
	"time"
)

var a string

func f() {
	fmt.Println(a)
}

func TestTiming(t *testing.T) {
	a = "hello world 1"
	go f()
	a = "hello world 2"
	go f()
	a = "hello world 3"
	time.Sleep(time.Millisecond)
}
