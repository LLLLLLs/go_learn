// Time        : 2019/07/08
// Description :

package _struct

import (
	"fmt"
	"golearn/leecode/linked_list/base"
	"reflect"
	"testing"
)

func TestEmpty(t *testing.T) {
	ea, eb := new(empty), new(empty)
	fmt.Printf("ea=%p,eb=%p\n", ea, eb)
	ea.hello()
	eb.world()
	fmt.Println(reflect.TypeOf(base.ListNode{}).Name())
}

type Foo struct {
	A string
	B *string
}

var f Foo

type Bar struct {
	value Foo
	ptr   *Foo
}

func newBar() Bar {
	return Bar{}
}

func TestFoo(t *testing.T) {
	fmt.Printf("%+v\n", f)
	bar := newBar()
	fmt.Printf("%+v\n", bar)
	fmt.Printf("%p\n", &bar.ptr)

	fmt.Printf("foo:%p,foo.A:%p\n", &f, &f.A)
}
