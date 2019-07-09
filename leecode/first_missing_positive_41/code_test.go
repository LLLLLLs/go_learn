// Time        : 2019/07/02
// Description :

package first_missing_positive_41

import (
	"fmt"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	fmt.Println(firstMissingPositive([]int{-1, 4, 2, 1, 9, 10}))
}
