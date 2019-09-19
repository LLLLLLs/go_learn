// Time        : 2019/09/10
// Description :

package loader

import (
	"fmt"
	"golearn/sundry/profile/example"
	"testing"
)

func BenchmarkIntIndexToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexToString(i)
	}
}

func BenchmarkInt64IndexToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexToString(int64(i))
	}
}

func BenchmarkUintIndexToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexToString(uint(i))
	}
}

func BenchmarkStringIndexToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexToString("111")
	}
}

func BenchmarkTypeIndexToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexToString(example.Index(123))
	}
}

func BenchmarkFloatIndexToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexToString((1.23))
	}
}

func BenchmarkLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100000; j++ {
			IndexToString(j)
			IndexToString(uint(j))
		}
	}
}

type Object struct {
	A string
	B string
}

func (o Object) Format(f fmt.State, c rune) {
	var err error
	switch c {
	case 's':
		str := fmt.Sprintf("{A:%s,B:%s}", o.A, o.B)
		_, err = f.Write([]byte(str))
	case 'v':
		str := fmt.Sprintf("{A:%s,B:%s}", o.A, o.B)
		_, err = f.Write([]byte(str))
	case 'b':
		str := fmt.Sprintf("{A:%s,B:%s}", o.A, o.B)
		_, err = f.Write([]byte(str))
	default:
		panic("wrong verb")

	}
	if err != nil {
		panic(err)
	}
}

func (o Object) String() string {
	return o.A
}

func TestPrint(t *testing.T) {
	obj := Object{
		A: "test",
		B: "string",
	}
	fmt.Println(obj)
	fmt.Printf("%+v\n", obj)
	fmt.Printf("%v\n", obj)
	fmt.Printf("%s\n", obj)
	fmt.Printf("%b\n", obj)
	fmt.Printf("%d\n", obj)
}
