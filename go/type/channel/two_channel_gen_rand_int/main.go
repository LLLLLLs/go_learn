// Time        : 2019/07/15
// Description : 实现两个 goroutine，其中一个产生随机数，另外一个读取数字并打印到标准输出。最终输出五个随机数

package main

import (
	"fmt"
	"golearn/util/randutil"
)

func main() {
	ch := make(chan int)
	sig := make(chan struct{})
	go func() {
		for i := 0; i < 5; i++ {
			ch <- randutil.RandInt(1, 5)
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-ch)
		}
		close(sig)
	}()
	<-sig
}
