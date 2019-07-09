// Time        : 2019/07/09
// Description :

package _struct

import (
	"fmt"
	"testing"
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
