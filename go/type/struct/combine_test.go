// Time        : 2019/08/19
// Description :

package _struct

import (
	"fmt"
	"golearn/utils"
	"testing"
)

type CombineA struct {
	M int
	N int
}

type CombineB struct {
	CombineA
	X int
	Y int
}

func TestCombine(t *testing.T) {
	b := CombineB{
		CombineA: CombineA{},
		X:        0,
		Y:        0,
	}
	fmt.Println(utils.MarshalToString(b))
}
