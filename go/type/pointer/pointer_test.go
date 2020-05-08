// Time        : 2019/01/18
// Description :

package pointer

import (
	"fmt"
	"golearn/go/type/pointer/private"
	"reflect"
	"testing"
	"unsafe"
)

func TestPtrBase(t *testing.T) {
	pointerBase()
}

func TestUnsafePtr(t *testing.T) {
	unsafePtr()
}

func TestUnsafe2Uint(t *testing.T) {
	unsafe2Uint()
}

func BenchmarkUnsafePtr(b *testing.B) {
	ts := struct {
		A int64
		B int64
	}{1, 2}
	for i := 0; i < b.N; i++ {
		_ = unsafe.Pointer(&ts)
	}
}

func BenchmarkReflectPtr(b *testing.B) {
	ts := struct {
		A int64
		B int64
	}{1, 2}
	for i := 0; i < b.N; i++ {
		_ = reflect.ValueOf(&ts).Pointer()
	}
}

func TestPointHeadOrTail(t *testing.T) {
	c := D{
		IA: &A{},
		IB: &B{},
		//IC: &C{},
	}
	value := reflect.ValueOf(c)
	for i := 0; i < value.NumField(); i++ {
		fmt.Println(value.Field(i).Type().Name(), reflect.ValueOf(value.Field(i).Interface()).Pointer())
	}
	fmt.Println(6704720 * 16 / 10)
	fmt.Println(reflect.ValueOf(c.IA).Pointer())
	fmt.Println(reflect.ValueOf(c.IB).Pointer())
	//fmt.Println(reflect.ValueOf(c.IC).Pointer())
	a := struct{}{}
	b := struct{}{}
	fmt.Println(a == b)
}

type ptrInfo struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

type s1 struct {
	A int
}

type s2 struct {
	B string
}

func TestPtr(t *testing.T) {
	var x = 124

	var y = "123"
	xp := (*ptrInfo)(unsafe.Pointer(&x))
	fmt.Println(xp)
	yp := (*ptrInfo)(unsafe.Pointer(&y))
	yp.typ = xp.typ
	yp.data = xp.data
	fmt.Printf("%+v\n", y)
}

func TestPtr2(t *testing.T) {
	s := s1{A: 1}
	//s1p := (*ptrInfo)(unsafe.Pointer(&s))

	ptr2 := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(s))
	//s2p := (*ptrInfo)(ptr2)
	//s2p.typ = s1p.typ
	s2d := (*s1)(ptr2)
	s2d.A = 100
	//fmt.Println(*ptr2)
	list := (*[2]s1)(unsafe.Pointer(&s))
	fmt.Println(*list)
}

// 修改结构体私有变量
func TestPrivate(t *testing.T) {
	a := private.NewA(1, "123")
	aa := (*int)(unsafe.Pointer(&a))
	*aa = 100
	ab := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + unsafe.Sizeof(int(0))))
	*ab = "hello"
	fmt.Printf("%+v\n", a)
}
