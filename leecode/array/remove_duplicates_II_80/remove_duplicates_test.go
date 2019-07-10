// Time        : 2019/07/10
// Description :

package remove_duplicates_II_80

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
