// Time        : 2019/07/31
// Description :

package find_min_in_rotated_sorted_array_II_154

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
// (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
//
// Find the minimum element.
//
// The array may contain duplicates.
//
// Example 1:
//
// Input: [1,3,5]
// Output: 1
// Example 2:
//
// Input: [2,2,2,0,1]
// Output: 0

func findMin(nums []int) int {
	var begin, end = 0, len(nums) - 1
	for begin < end && (nums[begin] == nums[begin+1] || nums[end] == nums[end-1]) {
		if nums[begin] == nums[begin+1] {
			begin++
		} else {
			end--
		}
	}
	if begin == end || nums[begin] < nums[end] {
		return nums[begin]
	}
	mid := (begin + end + 1) / 2
	if nums[mid] < nums[mid-1] {
		return nums[mid]
	}
	if nums[mid] < nums[begin] {
		return findMin(nums[begin:mid])
	} else {
		return findMin(nums[mid+1 : end+1])
	}
}
