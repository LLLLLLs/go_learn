// Time        : 2019/07/23
// Description :

package III_123

import (
	"fmt"
	"testing"
)

func TestBuyAndSell(t *testing.T) {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{1, 4, 2}))
}

func TestBuyAndSell2(t *testing.T) {
	fmt.Println(maxProfit2([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println(maxProfit2([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit2([]int{1, 4, 2}))
}
