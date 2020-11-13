//@author: lls
//@time: 2020/08/27
//@desc:

package franction_166

import (
	"fmt"
	"testing"
)

func TestFractionToDecimal(t *testing.T) {
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(3, 4))
	fmt.Println(fractionToDecimal(4, 9))
}
