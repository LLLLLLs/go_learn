// Time        : 2019/07/08
// Description : 测试方法receiver对性能影响

package _struct

import "fmt"

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
