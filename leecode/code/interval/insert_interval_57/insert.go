// Time        : 2019/07/04
// Description :

package insert_interval_57

import "sort"

// Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
//
// You may assume that the intervals were initially sorted according to their start times.
//
// Example 1:
//
// Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
// Output: [[1,5],[6,9]]
// Example 2:
//
// Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// Output: [[1,2],[3,10],[12,16]]
// Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{newInterval}
	for i := range intervals {
		tmp := result[len(result)-1]
		if intervals[i][1] < tmp[0] {
			result = append(result[:len(result)-1], intervals[i], tmp)
		} else if intervals[i][0] > tmp[1] {
			result = append(result, intervals[i:]...)
			break
		} else {
			tmp[0] = min(intervals[i][0], tmp[0])
			tmp[1] = max(intervals[i][1], tmp[1])
		}
	}
	return result
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
