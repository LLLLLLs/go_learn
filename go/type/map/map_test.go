// Time        : 2019/06/25
// Description :

package _map

import (
	"fmt"
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
