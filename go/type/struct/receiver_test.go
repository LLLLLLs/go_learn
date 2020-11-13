// Time        : 2019/07/08
// Description :

package _struct

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestReceiverPtr(t *testing.T) {
	t1 := Type1{}
	t2 := Type2{}
	fmt.Printf("%p\n", &t1)
	t1.Ptr()
	fmt.Printf("%p\n", &t2)
	t2.Ptr()
}

func TestStructWithMapOrPtr(t *testing.T) {
	s := sWithMap{m: map[int]int{}}
	fmt.Println(s)
	s.add(1)
	fmt.Println("s.add(1)", s)
	s.addPtr(2)
	fmt.Println("s.addPtr(2)", s)

	s2 := sWithPtr{m: &s}
	fmt.Println(s2, s2.m)
	s2.add(3)
	fmt.Println("a2.add(3)", s2, s2.m)

	for i := 0; i < 100; i++ {
		s.add(i)
		fmt.Println(s)
	}
}

func TestArray10(t *testing.T) {
	s := tWithArray10{array: [10]int{}}
	s2 := s
	s2.array[1] = 100
	fmt.Println(s)
	fmt.Println(s2)
	s.set(1, 1)
	fmt.Println(s)
}

func TestMap(t *testing.T) {
	var ori, cpy unsafe.Pointer
	vm := tWithMap{m: make(map[int]struct{}, 5)}
	for i := 0; i < 10; i++ {
		cpy = vm.value(i)
		ori = unsafe.Pointer(&vm.m)
		if cpy != ori {
			fmt.Printf("value %d, ori:%v, cpy:%v, %v\n", i, ori, cpy, vm.m)
		}
	}
	pm := tWithMap{m: make(map[int]struct{}, 5)}
	for i := 0; i < 10; i++ {
		cpy = pm.point(i)
		ori = unsafe.Pointer(&pm.m)
		if cpy != ori {
			fmt.Printf("value %d, ori:%v, cpy:%v, %v\n", i, ori, cpy, vm.m)
		}
	}
}

func TestSlice(t *testing.T) {
	vs := tWithSlice{s: make([]int, 0, 5)}
	prevP := unsafe.Pointer(&vs.s)
	nowP := prevP
	for i := 0; i < 10; i++ {
		nowP = vs.value(i)
		if prevP != nowP {
			fmt.Printf("value %d\n", i)
			fmt.Println(vs.s)
		}
		prevP = nowP
	}
	ps := tWithSlice{s: []int{}}
	prevP = unsafe.Pointer(&ps.s)
	nowP = prevP
	for i := 0; i < 10; i++ {
		nowP = ps.point(i)
		if prevP != nowP {
			fmt.Printf("point %d\n", i)
		}
		prevP = nowP
	}
}
