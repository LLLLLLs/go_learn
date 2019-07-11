// Time        : 2019/07/11
// Description :

package merge_sorted_88

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}
	merge(num1, 3, num2, len(num2))
	fmt.Println(num1)
}
