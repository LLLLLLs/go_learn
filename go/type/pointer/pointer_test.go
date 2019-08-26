// Time        : 2019/01/18
// Description :

package pointer

import (
	"fmt"
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
