// Time        : 2019/09/23
// Description :

package bit_operation

import (
	"fmt"
	"testing"
)

func TestBitOperation(t *testing.T) {
	a := uint8(0b10011110)
	b := uint8(0b11110101)
	and := a & b
	or := a | b
	xor := a ^ b
	not := ^a
	fmt.Printf("a = %b , b = %b\n", a, b)
	fmt.Printf("a and b = (a & b) %b\n", and)
	fmt.Printf("a or  b = (a | b) %b\n", or)
	fmt.Printf("a xor b = (a ^ b) %08b\n", xor)
	fmt.Printf("  not a = (  ^ a) %08b\n", not)
	// &^ = bit clear
	// 将b中为1的位对应a的位清零，a中其余位不变
	bc := a &^ b
	fmt.Printf("a bc  b = (a &^ b) %08b\n", bc)

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
