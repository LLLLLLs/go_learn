// Time        : 2019/09/23
// Description :

package bit_operation

import (
	"fmt"
	"testing"
)

type UnionPosition int16

const (
	Leader UnionPosition = 1 << iota
	DeputyLeader
	Elite
	Regular
)

func (pos UnionPosition) In(poss ...UnionPosition) bool {
	var mix UnionPosition
	for _, p := range poss {
		if mix&p == 1 {
			continue
		}
		mix += p
	}
	return pos&(mix) == 1
}

func TestIn(t *testing.T) {
	fmt.Println(Leader.In(Leader, DeputyLeader, Elite))
	fmt.Println(Leader.In(DeputyLeader, Elite))
}

func TestBitOperation(t *testing.T) {
	a := uint8(0b10011110)
	b := uint8(0b11110101)
	and := a & b
	or := a | b
	xor := a ^ b
	not := ^a
	fmt.Printf("a = %b , b = %b\n", a, b)
	fmt.Printf("a and b = (%b &  %b) = %b\n", a, b, and)
	fmt.Printf("a or  b = (%b |  %b) = %b\n", a, b, or)
	fmt.Printf("a xor b = (%b ^  %b) = %08b\n", a, b, xor)
	fmt.Printf("  not a = (         ^  %b) = %08b\n", a, not)
	// &^ = bit clear
	// 将b中为1的位对应a的位清零，a中其余位不变
	bc := a &^ b
	fmt.Printf("a bc  b = (%b &^ %b) = %08b\n", a, b, bc)

	fmt.Println()
	c := uint8(0b10000111)
	fmt.Printf("c = %08b\n", c)
	c = c >> 1
	fmt.Printf("c = c >> 1 => c = %08b\n", c)
	c = c << 2
	fmt.Printf("c = c << 2 => c = %08b\n", c)
	c >>= 1
	fmt.Printf("c >>= 1 => c = %08b\n", c)
}
