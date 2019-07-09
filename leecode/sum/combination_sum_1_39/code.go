// Time        : 2019/07/01
// Description :

package combination_sum_1_39

import "sort"

// Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
//
// The same repeated number may be chosen from candidates unlimited number of times.
//
// Note:
//
// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
// Example 1:
//
// Input: candidates = [2,3,6,7], target = 7,
// A solution set is:
// [
//   [7],
//   [2,2,3]
// ]
// Example 2:
//
// Input: candidates = [2,3,5], target = 8,
// A solution set is:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]

func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	result := make([][]int, 0)
	combine(candidates, []int{}, 0, target, &result)
	return result
}

func combine(candidates, tmp []int, midSum, target int, result *[][]int) {
	var mid = make([]int, len(tmp))
	copy(mid, tmp)
	if len(candidates) == 0 {
		return
	}
	for i := range candidates {
		midSum += candidates[i]
		if midSum == target {
			*result = append(*result, append(mid, candidates[i]))
			return
		} else if midSum < target {
			combine(candidates[i:], append(mid, candidates[i]), midSum, target, result)
		} else {
			return
		}
		midSum -= candidates[i]
	}
}
