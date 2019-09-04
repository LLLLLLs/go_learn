// Time        : 2019/09/03
// Description :

package typeconvert

import (
	"fmt"
	"testing"
)

func TestNumberToString(t *testing.T) {
	fmt.Println(NumberToString(100))
	fmt.Println(NumberToString(int8(100)))
	fmt.Println(NumberToString(int16(100)))
	fmt.Println(NumberToString(int64(100)))
	fmt.Println(NumberToString(uint8(100)))
	fmt.Println(NumberToString(uint16(100)))
	fmt.Println(NumberToString(uint64(100)))
	fmt.Println(NumberToString(float32(100.1)))
	fmt.Println(NumberToString(100.2222))
	fmt.Println(NumberToString(-1.11))
}
