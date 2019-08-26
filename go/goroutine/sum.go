// Time        : 2019/08/23
// Description :

package goroutine

import (
	"fmt"
	"time"
)

var (
	num               = 99999999
	sum3, sum5, sum15 = make(chan int64, 1), make(chan int64, 1), make(chan int64, 1)
)

func sumWithGoroutine() {
	start := time.Now()
	go Sum3()
	go Sum5()
	go Sum15()
	x3 := <-sum3
	x5 := <-sum5
	x15 := <-sum15
	goSum := x3 + x5 - x15
	fmt.Println("并行", goSum)
	fmt.Println(time.Now().Sub(start))
}

func sumWithoutGoroutine() {
	start := time.Now()
	x3 := Sum3()
	x5 := Sum5()
	x15 := Sum15()
	sum := x3 + x5 - x15
	fmt.Println("串行", sum)
	fmt.Println(time.Now().Sub(start))
}

func Sum3() int64 {
	var sum = int64(0)
	for i := 0; i < num; i++ {
		if i%3 == 0 {
			sum += int64(i)
		}
	}
	sum3 <- sum
	return sum
}

func Sum5() int64 {
	var sum = int64(0)
	for i := 0; i < num; i++ {
		if i%5 == 0 {
			sum += int64(i)
		}
	}
	sum5 <- sum
	return sum
}

func Sum15() int64 {
	var sum = int64(0)
	for i := 0; i < num; i++ {
		if i%15 == 0 {
			sum += int64(i)
		}
	}
	sum15 <- sum
	return sum
}
