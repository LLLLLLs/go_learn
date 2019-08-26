// Time        : 2019/07/31
// Description :

package find_min_in_rotated_sorted_array_153

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
// (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
//
// Find the minimum element.
//
// You may assume no duplicate exists in the array.
//
// Example 1:
//
// Input: [3,4,5,1,2]
// Output: 1
// Example 2:
//
// Input: [4,5,6,7,0,1,2]
// Output: 0

func findMin(nums []int) int {
	if len(nums) == 1 || nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	mid := len(nums) / 2
	if nums[mid] < nums[mid-1] {
		return nums[mid]
	}
	if nums[mid] < nums[0] {
		return findMin(nums[:mid])
	} else {
		return findMin(nums[mid+1:])
	}
}
