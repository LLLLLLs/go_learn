// Time        : 2019/08/28
// Description :

package calc_sss_percent

import (
	"fmt"
	"testing"
)

func TestCalcPercent(t *testing.T) {
	for i := 1; i < 10; i++ {
		fmt.Println(calcPercent(i))
	}
}

func TestTotal(t *testing.T) {
	for i := 1; i < 20; i++ {
		var percent float64
		for j := 1; j <= i; j++ {
			percent += calcPercent(j)
		}
		fmt.Println(i, 1+6*(i-1), percent)
	}
}
