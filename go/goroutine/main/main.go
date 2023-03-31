package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(10)
	a := int32(0)
	b := &a
	wg := sync.WaitGroup{}
	subFunc := func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			if a == *b {
				atomic.AddInt32(&a, 1)
			}
			// runtime.Gosched()
		}
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go subFunc()
	}
	wg.Wait()
	fmt.Println(a)
}
