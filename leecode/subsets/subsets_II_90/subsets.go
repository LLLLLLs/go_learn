// Time        : 2019/07/12
// Description :

package subsets_II_90

import "sort"

// Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).
//
// Note: The solution set must not contain duplicate subsets.
//
// Example:
//
// Input: [1,2,2]
// Output:
// [
//   [2],
//   [1],
//   [1,2,2],
//   [2,2],
//   [1,2],
//   []
// ]

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{{}}
	for k := 0; k <= len(nums); k++ {
		for i := range nums {
			if i == 0 || nums[i] != nums[i-1] {
				recursive(nums[i+1:], []int{nums[i]}, k, &result)
			}
		}
	}
	return result
}

func recursive(nums, mid []int, k int, result *[][]int) {
	if len(nums) < k {
		return
	}
	if k == 0 {
		*result = append(*result, append(mid))
		return
	}
	for i := range nums {
		if i == 0 || nums[i] != nums[i-1] {
			recursive(nums[i+1:], append(append([]int{}, mid...), nums[i]), k-1, result)
		}
	}
}
