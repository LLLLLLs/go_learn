// Time        : 2019/01/04
// Description :

package channel

import (
	"fmt"
	"time"
)

func chanTest() {
	c := make(chan int, 10)
	count := 10
	go read(c, count)
	go write(c, count)
	select {}
}

func read(c chan int, count int) {
	ticker := time.NewTicker(time.Second)
	for i := 0; i < count*2; i++ {
		<-ticker.C
		fmt.Println("receive:", <-c)
	}
}

func write(c chan int, count int) {
	ticker := time.NewTicker(time.Millisecond * 500)
	for i := 0; i < count; i++ {
		<-ticker.C
		c <- i
		fmt.Println("send:", i)
	}
	fmt.Println("close")
	close(c)
}
