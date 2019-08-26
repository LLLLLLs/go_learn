// Time        : 2019/07/31
// Description :

package find_peak_element_162

// A peak element is an element that is greater than its neighbors.
//
// Given an input array nums, where nums[i] ≠ nums[i+1], find a peak element and return its index.
//
// The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.
//
// You may imagine that nums[-1] = nums[n] = -∞.
//
// Example 1:
//
// Input: nums = [1,2,3,1]
// Output: 2
// Explanation: 3 is a peak element and your function should return the index number 2.
// Example 2:
//
// Input: nums = [1,2,1,3,5,6,4]
// Output: 1 or 5
// Explanation: Your function can return either index number 1 where the peak element is 2,
//              or index number 5 where the peak element is 6.
// Note:
// Your solution should be in logarithmic complexity.

// 思路：
// 	问题关键在于题目提到nums[-1]=nums[n]=-inf
// 	因此我们在比较下标mid和mid+1的值(mid = (left + right)/2)
//		若 mid < mid + 1 表示右半部存在升序序列，而最终却到达-inf，因此右半部存在peak;
//		若 mid > mid + 1 表示左半部存在降序序列，而起始值确实-inf，因此左半部存在peak。
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
