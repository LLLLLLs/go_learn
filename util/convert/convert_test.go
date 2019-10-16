// Time        : 2019/09/03
// Description :

package convert

import (
	"fmt"
	"testing"
)

func TestNumberToString(t *testing.T) {
	fmt.Println(ToString(100))
	fmt.Println(ToString(int8(100)))
	fmt.Println(ToString(int16(100)))
	fmt.Println(ToString(int64(100)))
	fmt.Println(ToString(uint8(100)))
	fmt.Println(ToString(uint16(100)))
	fmt.Println(ToString(uint64(100)))
	fmt.Println(ToString(float32(100.1)))
	fmt.Println(ToString(100.2222))
	fmt.Println(ToString(-1.11))
}
