// @author: lls
// @date: 2021/8/18
// @desc:

package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestWG(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go hello(wg)
	wg.Wait()
}

func hello(wg *sync.WaitGroup) {
	fmt.Println("hello")
	wg.Done()
}
