//@time:2020/01/20
//@desc:

package util

import (
	"fmt"
	"runtime"
)

func WithRecover(f func()) {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 1024)
			n := runtime.Stack(buf, false)
			fmt.Printf("panic Error: %v;stack: %s\n", e, buf[:n])
		}
	}()
	f()
}
