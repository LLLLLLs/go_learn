// Time        : 2019/11/04
// Description :

package int

import (
	"fmt"
	"testing"
)

func TestSignedOverflow(t *testing.T) {
	a, b := int8(0b01111111), int8(1)
	fmt.Println((a + b) >> 1)          // -64
	fmt.Println(int8(uint8(a+b) >> 1)) // 64
	fmt.Println((a + b) / 2)           // -64
	fmt.Println(int8(uint8(a+b) / 2))  // 64
}
