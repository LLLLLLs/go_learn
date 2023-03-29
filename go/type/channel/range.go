package channel

import (
	"fmt"
)

func ForSelect(c chan int, identity int) {
	for {
		select {
		case data := <-c:
			fmt.Println(identity, ":", data)
		}
	}
}

func Range(c chan int, identity int) {
	for data := range c {
		fmt.Println(identity, ":", data)
	}
}

func Send(c chan int) {
	for i := 0; ; i++ {
		c <- i
		// time.Sleep(time.Second)
	}
}
