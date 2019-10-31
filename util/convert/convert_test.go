// Time        : 2019/09/03
// Description :

package convert

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestNumberToString(t *testing.T) {
	fmt.Println(ToString(100))
	fmt.Println(ToString(int8(100)))
	fmt.Println(ToString(int16(100)))
	fmt.Println(ToString(int64(100)))
	fmt.Println(ToString(uint8(100)))
	fmt.Println(ToString(uint16(100)))
	fmt.Println(ToString(uint64(100)))
	fmt.Println(ToString(float32(100.1)))
	fmt.Println(ToString(100.2222))
	fmt.Println(ToString(-1.11))
}

func BenchmarkSprintf_Int64(b *testing.B) {
	var x = int64(1)
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%v", x)
	}
}

func BenchmarkFormatIntAndReflect_Int64(b *testing.B) {
	var x = int64(1)
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatInt(reflect.ValueOf(x).Int(), 10)
	}
}

func BenchmarkSprintf_G_Float32(b *testing.B) {
	var x = float32(1.11)
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%g", x)
	}
}

func BenchmarkSprintf_V_Float32(b *testing.B) {
	var x = float32(1.11)
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%v", x)
	}
}

func BenchmarkSprintf_Float64(b *testing.B) {
	var x = float64(1)
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%g", x)
	}
}

func BenchmarkFormatIntAndReflect_Float64(b *testing.B) {
	var x = float64(1)
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatFloat(x, 'g', -1, 64)
	}
}
