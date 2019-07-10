// Time        : 2019/07/10
// Description :

package subsets_78

// Given a set of distinct integers, nums, return all possible subsets (the power set).
//
// Note: The solution set must not contain duplicate subsets.
//
// Example:
//
// Input: nums = [1,2,3]
// Output:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	for i := 0; i <= len(nums); i++ {
		mid := make([]int, 0, i)
		recursive(nums, mid, i, &result)
	}
	return result
}

func recursive(nums, mid []int, k int, result *[][]int) {
	if len(nums) < k {
		return
	}
	if k == 0 {
		*result = append(*result, mid)
	}
	for i := range nums {
		tmp := make([]int, len(mid), cap(mid))
		copy(tmp, mid)
		tmp = append(tmp, nums[i])
		recursive(nums[i+1:], tmp, k-1, result)
	}
}

// 用二进制来表示是否是子集
// 000 = 空集
// 001 = [3]
// 110 = [1,2]
// 结果 = 000,001,010...111
func subsetsBinary(nums []int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < 1<<uint(len(nums)); i++ {
		tmp := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if i&(1<<uint(j)) != 0 {
				tmp = append(tmp, nums[len(nums)-1-j])
			}
		}
		result = append(result, tmp)
	}
	return result
}
