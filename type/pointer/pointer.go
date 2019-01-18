// Time        : 2019/01/18
// Description :

package pointer

import (
	"fmt"
	"reflect"
	"unsafe"
)

func pointerBase() {
	ts := struct {
		A int64
		B int64
	}{1, 2}
	fmt.Printf("%p,%p\n", &ts.A, &ts.B)
}

// 使用unsafe.Pointer效率比reflect高数十倍
func unsafePtr() {
	ts := struct {
		A int64
		B int64
	}{1, 2}
	ptr := unsafe.Pointer(&ts)
	fmt.Printf("ptr:%p,unsafe:%v,reflect:%x\n", &ts, ptr, reflect.ValueOf(&ts).Pointer())
}

func unsafe2Uint() {
	ts := struct {
		A int64
		B int64
	}{1, 2}
	var uintPtr uintptr
	unsafePtr := unsafe.Pointer(&ts)
	uintPtr = uintptr(unsafePtr)
	fmt.Printf("%b\t%o\t%d\t%x\n", uintPtr, uintPtr, uintPtr, uintPtr)
}
