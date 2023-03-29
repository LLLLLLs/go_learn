package _type

import (
	"fmt"
	"testing"
)

func TestTypeAssert(t *testing.T) {
	var raw interface{} = int32(0)
	switch v := raw.(type) {
	case Int:
		fmt.Println(int(v))
	}
}

type Int interface {
	int | int32
}
