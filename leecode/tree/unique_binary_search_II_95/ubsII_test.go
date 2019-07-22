// Time        : 2019/07/18
// Description :

package unique_binary_search_II_95

import (
	"fmt"
	"testing"
)

func TestUbs(t *testing.T) {
	trees := generateTrees(10)
	for i := range trees {
		fmt.Println(trees[i].PreorderTraversal())
	}
}

func TestSliceEqual(t *testing.T) {
	l1 := [2]int{0, 1}
	l2 := [2]int{0, 1}
	fmt.Println(l1 == l2)
}
