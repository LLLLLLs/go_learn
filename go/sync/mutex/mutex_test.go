//@author: lls
//@time: 2020/05/20
//@desc:

package mutex

import (
	"sync"
	"testing"
	"time"
)

func TestDeadlock(t *testing.T) {
	a, b := &sync.Mutex{}, &sync.Mutex{}
	go deadlock(a, b)
	go deadlock(b, a)
	time.Sleep(time.Minute)
}
