// Time        : 2019/07/04
// Description :

package merge_intervals_56

import "sort"

// Given a collection of intervals, merge all overlapping intervals.
//
// Example 1:
//
// Input: [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
// Example 2:
//
// Input: [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.
// NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		tmp := result[len(result)-1]
		if intervals[i][0] <= tmp[1] {
			tmp[1] = max(intervals[i][1], tmp[1])
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}
