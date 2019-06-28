// Time        : 2019/06/28
// Description :

package search_rotated_sorted_array_33

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{2}, 0))
	fmt.Println(search([]int{2, 3, 4, 5, 6, 7, 8, 9, 1}, 9))
	fmt.Println(search([]int{8, 9, 2, 3, 4}, 9))
}
