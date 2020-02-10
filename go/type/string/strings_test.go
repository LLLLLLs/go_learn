// Time        : 2019/11/08
// Description :

package string

import (
	"fmt"
	"golearn/util"
	"strings"
	"testing"
)

func TestFields(t *testing.T) {
	str := "hello world"
	fmt.Println(strings.Fields(str))
	str = "hello      wor ld"
	fmt.Println(strings.Fields(str))
}

func TestMap(t *testing.T) {
	str := "ABCDE"
	result := strings.Map(func(r rune) rune {
		return r + 1
	}, str)
	fmt.Println(str, result)
}

func TestMarshal(t *testing.T) {
	var data = struct {
		A string
		B string
	}{
		A: "Hello",
		B: "world",
	}
	var str = util.MarshalToString(data)
	for i := 0; i < 10; i++ {
		fmt.Println(str)
		str = util.MarshalToString(str)
	}
}

const (
	A int = iota
	B
	C
	D
	E int = 5
	F
	G
	H
	I int = iota + 20
	J
	K
)

func TestIota(t *testing.T) {
	fmt.Println(A, B, C, D, E, F, G, H, I, J, K)
}
