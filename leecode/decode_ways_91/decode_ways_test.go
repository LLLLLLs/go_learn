// Time        : 2019/07/12
// Description :

package decode_ways_91

import (
	"fmt"
	"testing"
)

func TestDecodeWays(t *testing.T) {
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("2262"))
	fmt.Println(numDecodings("22626"))
	fmt.Println(numDecodings("226262"))
	fmt.Println(numDecodings("2262626"))
	fmt.Println(numDecodings("22626262"))
	fmt.Println(numDecodings("226262622"))
	fmt.Println(numDecodings("2262626222"))
	fmt.Println(numDecodings("0"))
	fmt.Println(numDecodings("00"))
	fmt.Println(numDecodings("100"))
}
