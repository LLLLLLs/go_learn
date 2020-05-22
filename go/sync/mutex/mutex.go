//@author: lls
//@time: 2020/05/20
//@desc:

package mutex

import (
	"fmt"
	"sync"
	"time"
)

func deadlock(a, b *sync.Mutex) {
	a.Lock()
	fmt.Println("lock 1")
	time.Sleep(time.Second)
	b.Lock()
	fmt.Println("lock 2")
	fmt.Println("done")
}
