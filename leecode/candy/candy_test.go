// Time        : 2019/07/25
// Description :

package candy

import (
	"fmt"
	"testing"
)

func TestCandy(t *testing.T) {
	fmt.Println(candy([]int{1, 0, 2}))
	fmt.Println(candy2([]int{1, 2, 87, 87, 87, 2, 1}))
	fmt.Println(candy2([]int{1, 3, 2, 2, 1}))
	fmt.Println(candy3([]int{1, 3, 2, 2, 1}))
}
