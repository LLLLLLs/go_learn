// Time        : 2019/07/10
// Description :

package search_rotated_sorted_array_II_81

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	fmt.Println(search(nums, 3))
	fmt.Println(search(nums, 2))
	fmt.Println(search(nums, 1))
	nums = []int{1, 3}
	fmt.Println(search(nums, 2))
	fmt.Println(search(nums, 3))
	fmt.Println(search(nums, 1))
	nums = []int{3, 1}
	fmt.Println(search(nums, 2))
	fmt.Println(search(nums, 3))
	fmt.Println(search(nums, 1))
	nums = []int{1, 1}
	fmt.Println(search(nums, 0))
	fmt.Println(search(nums, 3))
	fmt.Println(search(nums, 1))
}
