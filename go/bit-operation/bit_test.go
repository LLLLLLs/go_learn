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

	c := 0b00000111
	fmt.Println(c)
}
