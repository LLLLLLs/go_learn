//@author: lls
//@time: 2020/08/03
//@desc:

package _struct

import "testing"

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

func BenchmarkTArrayValueReceiver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := tWithArray51024{array: [5120]int64{}}
		a := t.valueReceiver()
		_ = a
	}
}

func BenchmarkTArrayPointReceiver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := tWithArray51024{array: [5120]int64{}}
		a := t.pointReceiver()
		_ = a
	}
}

func BenchmarkTArrayValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := tWithArray51024{array: [5120]int64{}}
		ct := t.value()
		_ = ct
		//ct.array[0] = 100
		//t.array[1] = 200
	}
}

func BenchmarkTArrayPoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := tWithArray51024{array: [5120]int64{}}
		ct := t.point()
		_ = ct
		//ct.array[0] = 1000
		//t.array[1] = 2000
	}
}
