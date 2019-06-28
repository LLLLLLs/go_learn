// Time        : 2019/06/28
// Description :

package next_permutation_31

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	printList([]int{1, 2, 3})
	printList([]int{1, 1, 5})
	printList([]int{3, 2, 1})
}

func printList(list []int) {
	nextPermutation(list)
	fmt.Println(list)
}
