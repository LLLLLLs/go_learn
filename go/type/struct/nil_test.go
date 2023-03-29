package _struct

import (
	"fmt"
	"testing"
)

func TestNil(t *testing.T) {
	// getNil().PrintA()
	var n interface{} = nil
	_, ok := n.(*nilType)
	fmt.Println(ok)
}

func getNil() *nilType {
	return nil
}
