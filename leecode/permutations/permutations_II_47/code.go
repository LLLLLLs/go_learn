// Time        : 2019/07/03
// Description :

package permutations_II_47

// Given a collection of numbers that might contain duplicates, return all possible unique permutations.
//
// Example:
//
// Input: [1,1,2]
// Output:
// [
//   [1,1,2],
//   [1,2,1],
//   [2,1,1]
// ]

func permuteUnique(nums []int) [][]int {
	var result = make([][]int, 0)
	recursive(nums, 0, &result)
	return result
}

func recursive(nums []int, index int, result *[][]int) {
	if index == len(nums)-1 {
		*result = append(*result, append([]int{}, nums...))
	}
	repeat := make(map[int]bool)
	for i := index; i < len(nums); i++ {
		if repeat[nums[i]] {
			continue
		}
		nums[index], nums[i] = nums[i], nums[index]
		recursive(nums, index+1, result)
		nums[index], nums[i] = nums[i], nums[index]
		repeat[nums[i]] = true
	}
}
