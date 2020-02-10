// Time        : 2019/07/09
// Description :

package _struct

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestLength(t *testing.T) {
	fmt.Println(unsafe.Sizeof(agg))
	fmt.Println(agg.field1 == nil)
	s1 := make([]foo, 0)
	s2 := make([]bar, 1)
	s3 := make([]bar, 2)
	s4 := make([]empty, 1)
	s5 := [2]bar{}
	fmt.Println(unsafe.Sizeof(s1))
	fmt.Println(unsafe.Sizeof(s2))
	fmt.Println(unsafe.Sizeof(s3))
	fmt.Println(unsafe.Sizeof(s4))
	fmt.Println(unsafe.Sizeof(s5))
}

func TestMapLength(t *testing.T) {
	var m = make(map[int]int)
	fmt.Println("空map[int]int:", unsafe.Sizeof(m))
	m[1] = 1
	fmt.Println("添加一个元素m[1]=1:", unsafe.Sizeof(m))
	var elem = Aggregation{}
	fmt.Println(unsafe.Sizeof(elem))
	var ptr = &elem
	fmt.Println(unsafe.Sizeof(ptr))
	var list = []int{1, 2, 3}
	fmt.Println(unsafe.Sizeof(list))

	fmt.Println(time.Unix(1576684453, 0))
}
