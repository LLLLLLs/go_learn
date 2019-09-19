// Time        : 2019/09/04
// Description :

package big

import (
	"fmt"
	"math/big"
	"testing"
)

func TestFloat(t *testing.T) {
	a := 143.66
	b := 14.55
	c := a - b
	fmt.Printf("a = (%T)%v\n", a, a)
	fmt.Printf("b = (%T)%v\n", b, b)
	fmt.Printf("a - b = (%T)%v\n", c, c)
	fmt.Println()
	d := 1129.6
	fmt.Printf("a = (%T)%v\n", d, d)
	fmt.Printf("d * 100 = (%T)%v\n", d*100, d*100)
}

// big 也无法消除十进制和二进制小数之间转换的精度问题
func TestBig(t *testing.T) {
	a := big.NewFloat(143.66)
	b := big.NewFloat(14.55)
	var c = new(big.Float)
	c.Sub(a, big.NewFloat(14.55))
	cf, acc := c.Float64()
	fmt.Printf("a = (%T)%v\n", a, a)
	fmt.Printf("b = (%T)%v\n", b, b)
	fmt.Printf("a - b = (%T)%v\n", cf, cf)
	fmt.Println(acc)
}
