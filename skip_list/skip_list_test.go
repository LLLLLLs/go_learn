// Time        : 2019/01/08
// Description :

package skip_list

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSkipList(t *testing.T) {
	rand.Seed(time.Now().Unix())
	l := newList()
	l.insert(1, 1)
	l.insert(2, 1)
	l.insert(3, 1)
	l.insert(4, 1)
	l.insert(5, 1)
	l.insert(7, 1)
	l.insert(8, 1)
	l.insert(10, 1)
	l.insert(11, 1)
	l.insert(15, 1)
	l.insert(14, 1)
	l.print()
	l.delete(14)
	fmt.Println()
	l.print()
	l.insert(14, 14)
	fmt.Println()
	l.print()
	fmt.Println("search 1", l.search(1).key)
	fmt.Println("search 9", l.search(9))
	fmt.Println("search 14", l.search(14))
	fmt.Println("search 15", l.search(15).key)
}
