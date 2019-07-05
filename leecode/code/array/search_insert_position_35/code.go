// Time        : 2019/06/28
// Description :

package search_insert_position_35

// Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
// You may assume no duplicates in the array.
//
// Example 1:
// Input: [1,3,5,6], 5
// Output: 2
//
// Example 2:
// Input: [1,3,5,6], 2
// Output: 1
//
// Example 3:
// Input: [1,3,5,6], 7
// Output: 4
//
// Example 4:
// Input: [1,3,5,6], 0
// Output: 0

func searchInsert(nums []int, target int) int {
	return searchRecursion(nums, target, 0, len(nums)-1)
}

func searchRecursion(nums []int, target, i, j int) int {
	length := len(nums)
	if length == 0 {
		return i
	}
	if length == 1 {
		if target > nums[0] {
			return j + 1
		}
		return i
	}
	midIndex := length / 2
	mid := nums[midIndex]
	if mid == target {
		return midIndex + i
	} else if target < mid {
		return searchRecursion(nums[:midIndex], target, i, i+midIndex-1)
	} else {
		return searchRecursion(nums[midIndex:], target, i+midIndex, j)
	}
}
