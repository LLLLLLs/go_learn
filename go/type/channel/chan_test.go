// Time        : 2019/01/04
// Description :

package channel

import (
	"fmt"
	"sync"
	"testing"
)

func TestChan(t *testing.T) {
	chanTest()
}

func TestCloseSig(t *testing.T) {
	wg := sync.WaitGroup{}
	closeSig := make(chan struct{})
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			var count = 0
			for {
				select {
				case <-closeSig:
					fmt.Println("close")
					wg.Done()
					return
				default:
					count++
					fmt.Println(count)
					if count == 5 {
						close(closeSig)
					}
				}
			}
		}()
	}
	wg.Wait()
}

func TestClose(t *testing.T) {
	c := make(chan int, 3)
	go produce(c)
	consume(c)
}

func produce(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func consume(c chan int) {
	for {
		i, ok := <-c
		if !ok {
			break
		}
		fmt.Println(i)
	}
}
