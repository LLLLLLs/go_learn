package script

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestB_Nil(t *testing.T) {
	a := new(A)
	a.B.Func()
	a.B.Empty()
}

type AA struct{}

func (*AA) Print() {
	fmt.Println("aa")
}

type BB struct{}

func (*BB) Print() {
	fmt.Println("bb")
}

func TestNil(t *testing.T) {
	aa := new(AA)
	bb := new(BB)
	cc := new(AA)
	fmt.Println(unsafe.Pointer(aa) == unsafe.Pointer(bb))
	fmt.Println(unsafe.Pointer(cc) == unsafe.Pointer(bb))
	reflect.NewAt(reflect.TypeOf(&AA{}), unsafe.Pointer(bb)).Elem().Interface().(*AA).Print()
	reflect.NewAt(reflect.TypeOf(&BB{}), unsafe.Pointer(aa)).Elem().Interface().(*BB).Print()
	fmt.Println(unsafe.Pointer(aa))

	list := []int32{1, 2, 3}
	list = list[:len(list)-1]
	fmt.Println(list)

	for i := 0; i < 1; i++ {
		fmt.Println(i)
	}
}
