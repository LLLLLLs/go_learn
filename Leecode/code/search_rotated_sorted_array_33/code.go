// Time        : 2019/06/28
// Description :

package search_rotated_sorted_array_33

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
// (例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2])。
//
// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//
// 你可以假设数组中不存在重复的元素。
//
// 你的算法时间复杂度必须是 O(log n) 级别。
//

func search(nums []int, target int) int {
	return searchRecursion(nums, target, 0, len(nums)-1)
}

func searchRecursion(nums []int, target, i, j int) int {
	length := len(nums)
	if length == 0 || (nums[length-1] > nums[0] && (target < nums[0] || target > nums[length-1])) {
		return -1
	}
	if length == 1 {
		if nums[0] != target {
			return -1
		}
	}
	if nums[0] == target {
		return i
	}
	if nums[length-1] == target {
		return j
	}
	mid := nums[length/2]
	if mid == target {
		return i + length/2
	}

	if (target > nums[0] && target < mid && mid > nums[0]) ||
		(target < nums[0] && target < mid && mid < nums[0]) ||
		(target > nums[0] && mid < nums[0]) {
		return searchRecursion(nums[:length/2], target, i, i+length/2-1)
	} else {
		return searchRecursion(nums[length/2:], target, i+length/2, j)
	}
}
