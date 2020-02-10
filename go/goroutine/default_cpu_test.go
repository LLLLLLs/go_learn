// Time        : 2019/12/24
// Description :

package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func printNum(n int) {
	for i := n * 3; i < (n+1)*3; i++ {
		fmt.Println(i)
	}
}

func TestDefaultCPUNum(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			printNum(n)
		}(i)
	}
	wg.Wait()
}

func TestSetCPU_1(t *testing.T) {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			printNum(n)
		}(i)
	}
	wg.Wait()
}

func TestSetCPU_Full(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			printNum(n)
		}(i)
	}
	wg.Wait()
}
