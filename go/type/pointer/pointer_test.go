// Time        : 2019/01/18
// Description :

package pointer

import (
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
