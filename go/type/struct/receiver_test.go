// Time        : 2019/07/08
// Description :

package _struct

import (
	"fmt"
	"testing"
)

func BenchmarkValue(b *testing.B) {
	str := ""
	for i := 0; i < 1000; i++ {
		str += "hello "
	}
	s := Type1{
		a: 1,
		b: 2,
		c: 100.0,
		d: str,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.A()
		s.B()
		s.C()
		s.D()
		s.Sum()
	}
}

func BenchmarkPointer(b *testing.B) {
	str := ""
	for i := 0; i < 1000; i++ {
		str += "hello "
	}
	s := Type2{
		a: 1,
		b: 2,
		c: 100.0,
		d: str,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.A()
		s.B()
		s.C()
		s.D()
		s.Sum()
	}
}

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
