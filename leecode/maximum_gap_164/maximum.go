// Time        : 2019/07/31
// Description :

package maximum_gap_164

import "math"

//Given an unsorted array, find the maximum difference between the successive elements in its sorted form.
//
//Return 0 if the array contains less than 2 elements.
//
//Example 1:
//
//Input: [3,6,9,1]
//Output: 3
//Explanation: The sorted form of the array is [1,3,6,9], either
//             (3,6) or (6,9) has the maximum difference 3.
//Example 2:
//
//Input: [10]
//Output: 0
//Explanation: The array contains less than 2 elements, therefore return 0.
//Note:
//
//You may assume all elements in the array are non-negative integers and fit in the 32-bit signed integer range.
//Try to solve it in linear time/space.

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var minimum, maximum = math.MaxInt32, math.MinInt32
	for i := range nums {
		if nums[i] < minimum {
			minimum = nums[i]
		}
		if nums[i] > maximum {
			maximum = nums[i]
		}
	}
	size := (maximum + 1 - minimum) / (len(nums) - 1)
	bucketsNum := int(math.Ceil(float64(maximum+1-minimum) / float64(size)))
	if size == 0 {
		bucketsNum = 1
	}
	buckets := make([][]int, bucketsNum)
	for i := range nums {
		var index int
		if size != 0 {
			index = (nums[i] - minimum) / size
		}
		if buckets[index] == nil {
			buckets[index] = []int{nums[i], nums[i]}
		} else {
			buckets[index][0] = min(buckets[index][0], nums[i])
			buckets[index][1] = max(buckets[index][1], nums[i])
		}
	}
	var maxGap, prev int
	first := true
	for i := range buckets {
		if buckets[i] == nil {
			continue
		}
		for j := range buckets[i] {
			if first {
				prev = buckets[i][j]
				first = false
				continue
			}
			maxGap = max(maxGap, buckets[i][j]-prev)
			prev = buckets[i][j]
		}
	}
	return maxGap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
