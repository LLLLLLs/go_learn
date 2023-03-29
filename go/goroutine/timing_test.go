// Time        : 2019/12/02
// Description :

package goroutine

import (
	"fmt"
	"sync"
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

func TestGoroutine(t *testing.T) {
	a := int32(0)
	wg := sync.WaitGroup{}
	wg.Add(2)
	subFunc := func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			if a == a {
				a++
			}
		}
	}
	go subFunc()
	go subFunc()
	wg.Wait()
	fmt.Println(a)
}
