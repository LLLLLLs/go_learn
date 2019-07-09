// Time        : 2019/07/03
// Description :

package permutations_I_46

// Given a collection of distinct integers, return all possible permutations.
//
// Example:
//
// Input: [1,2,3]
// Output:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

func permute(nums []int) [][]int {
	var result = make([][]int, 0)
	recursive(nums, 0, &result)
	return result
}

func recursive(nums []int, index int, result *[][]int) {
	if index == len(nums)-1 {
		*result = append(*result, append([]int{}, nums...))
	}
	for i := index; i < len(nums); i++ {
		nums[index], nums[i] = nums[i], nums[index]
		recursive(nums, index+1, result)
		nums[index], nums[i] = nums[i], nums[index]
	}
}
