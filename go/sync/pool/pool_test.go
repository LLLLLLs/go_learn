// Time        : 2019/11/01
// Description :

package pool

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	i := intPool.Get().(int)
	fmt.Println(i)
	i = intPool.Get().(int)
	fmt.Println(i)
	intPool.Put(321)
	i = intPool.Get().(int)
	fmt.Println(i)
	i = intPool.Get().(int)
	fmt.Println(i)
	intPool.Put(333)
	intPool.Put(222)
	i = intPool.Get().(int)
	fmt.Println(i)
	i = intPool.Get().(int)
	fmt.Println(i)
}

func TestSlicePool(t *testing.T) {
	s := slicePool.Get().([]int)
	s[0] = 10
	slicePool.Put(s)
	s = slicePool.Get().([]int)
	fmt.Println(s)
}
