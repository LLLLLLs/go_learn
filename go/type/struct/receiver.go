// Time        : 2019/07/08
// Description : 测试方法receiver对性能影响

package _struct

import (
	"fmt"
	"unsafe"
)

type Type1 struct {
	a int
	b int
	c float64
	d string
}

func (t Type1) A() int {
	return t.a
}

func (t Type1) B() int {
	return t.b
}

func (t Type1) C() float64 {
	return t.c
}

func (t Type1) D() string {
	return t.d
}

func (t Type1) Sum() float64 {
	return t.c + float64(t.a) + float64(t.b)
}

func (t Type1) Ptr() {
	fmt.Printf("%p\n", &t)
}

type Type2 struct {
	a int
	b int
	c float64
	d string
}

func (t *Type2) A() int {
	return t.a
}

func (t *Type2) B() int {
	return t.b
}

func (t *Type2) C() float64 {
	return t.c
}

func (t *Type2) D() string {
	return t.d
}

func (t *Type2) Sum() float64 {
	return t.c + float64(t.a) + float64(t.b)
}

func (t *Type2) Ptr() {
	fmt.Printf("%p\n", t)
}

type sWithMap struct {
	m map[int]int
}

func (s sWithMap) add(key int) {
	s.m[key]++
}

func (s *sWithMap) addPtr(key int) {
	s.m[key]++
}

type sWithPtr struct {
	m *sWithMap
}

func (s sWithPtr) add(key int) {
	s.m.addPtr(key)
}

type tWithArray10 struct {
	array [10]int
}

func (t tWithArray10) set(index, value int) {
	t.array[index] = value
}

func (t tWithArray51024) valueReceiver() int64 {
	return t.h
}

func (t *tWithArray51024) pointReceiver() int64 {
	return t.h
}

type tWithArray51024 struct {
	array  [5 * 1024]int64
	gg     bool
	h      int64
	hjj    byte
	a      int
	array2 [5 * 1024]int64
}

func (t tWithArray51024) value() tWithArray51024 {
	return t
}

func (t *tWithArray51024) point() *tWithArray51024 {
	return t
}

type tWithMap struct {
	m map[int]struct{}
}

func (t tWithMap) value(i int) unsafe.Pointer {
	t.m[i] = struct{}{}
	return unsafe.Pointer(&t.m)
}

func (t *tWithMap) point(i int) unsafe.Pointer {
	t.m[i] = struct{}{}
	return unsafe.Pointer(&t.m)
}

type tWithSlice struct {
	s []int
}

func (t tWithSlice) value(n int) unsafe.Pointer {
	t.s = append(t.s, n)
	return unsafe.Pointer(&t.s)
}

func (t *tWithSlice) point(n int) unsafe.Pointer {
	t.s = append(t.s, n)
	return unsafe.Pointer(&t.s)
}
