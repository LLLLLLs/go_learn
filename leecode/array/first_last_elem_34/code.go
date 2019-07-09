// Time        : 2019/06/28
// Description :

package first_last_elem_34

// Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
//
// Your algorithm's runtime complexity must be in the order of O(log n).
//
// If the target is not found in the array, return [-1, -1].
//
// Example 1:
//
// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]
// Example 2:
//
// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]

func searchRange(nums []int, target int) []int {
	return searchRecursion(nums, target, 0, len(nums)-1)
}

func searchRecursion(nums []int, target, i, j int) []int {
	var result []int
	length := len(nums)
	if length <= 2 {
		result = []int{-1, -1}
		for k := range nums {
			if nums[k] == target {
				if result[0] == -1 {
					result[0] = i + k
				}
				result[1] = i + k
			}
		}
		return result
	}
	midIndex := length / 2
	mid := nums[midIndex]
	if mid == target {
		result = []int{midIndex + i, midIndex + i}
		var left, right = true, true
		for k := 1; left || right; k++ {
			if left && (midIndex-k < 0 || nums[midIndex-k] != target) {
				left = false
			} else if left {
				result[0] = midIndex + i - k
			}
			if right && (midIndex+k > len(nums)-1 || nums[midIndex+k] != target) {
				right = false
			} else if right {
				result[1] = midIndex + i + k
			}
		}
		return result
	} else if mid > target {
		return searchRecursion(nums[:midIndex], target, i, midIndex+i-1)
	} else {
		return searchRecursion(nums[midIndex:], target, midIndex+i, j)
	}
}
