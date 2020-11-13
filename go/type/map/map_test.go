// Time        : 2019/06/25
// Description :

package _map

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	var m = make(map[int16]bool)
	m[1] = true
	m[2] = true
	fmt.Println(len(m))
	delete(m, 1)
	fmt.Println(len(m))
	delete(m, 2)
	fmt.Println(len(m))
}

func TestCopyMap(t *testing.T) {
	m1 := map[string]bool{
		"1": true,
		"2": false,
	}
	m2 := m1
	delete(m2, "1")
	fmt.Println(m1)
}

func TestFuncMap(t *testing.T) {
	f := func(m map[int]bool) {
		delete(m, 1)
	}
	m := map[int]bool{1: true}
	f(m)
	fmt.Println(m)
}

func TestSyncMap(t *testing.T) {
	sm := sync.Map{}
	sm.Store(1, 1)
	sm.Store("abc", "cba")
	sm.Store(nil, nil)
	fmt.Println(sm.Load(nil))
	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	fmt.Printf("%065b\n", ^uintptr(0))
}

func TestMapExpand(t *testing.T) {
	m := make(map[int]int, 0)
	for i := 0; i < 100; i++ {
		m[i] = i
		fmt.Printf("%p\t%v\n", m, m)
	}
}

func TestMapComplex(t *testing.T) {
	m := make(map[complex64]int)
	m[1+1i] = 100
	fmt.Println(m[1+1i])
	x := complex64(1 + 2i)
	fmt.Println(real(x), imag(x))
	a, b := 1, 2
	x = complex(float32(a), float32(b))
	fmt.Println(x)
	fmt.Println(real(x), imag(x))
}

type s struct {
	A int
}

func TestModify(t *testing.T) {
	m := make(map[int]s)
	for i := 0; i < 10; i++ {
		m[i] = s{A: i}
	}
	fmt.Println(m)
	for k, v := range m {
		v.A += k
	}
	fmt.Println(m)
	for k, v := range m {
		v.A += k
		m[k] = v
	}
	fmt.Println(m)
	fmt.Println(m)
}
