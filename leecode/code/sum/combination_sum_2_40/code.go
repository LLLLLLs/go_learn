// Time        : 2019/07/02
// Description :

package combination_sum_2_40

import "sort"

// Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
//
// Each number in candidates may only be used once in the combination.
//
// Note:
//
// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
// Example 1:
//
// Input: candidates = [10,1,2,7,6,1,5], target = 8,
// A solution set is:
// [
//   [1, 7],
//   [1, 2, 5],
//   [2, 6],
//   [1, 1, 6]
// ]
// Example 2:
//
// Input: candidates = [2,5,2,1,2], target = 5,
// A solution set is:
// [
//   [1,2,2],
//   [5]
// ]

func combinationSum2(candidates []int, target int) [][]int {
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
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		midSum += candidates[i]
		if midSum == target {
			*result = append(*result, append(mid, candidates[i]))
			return
		} else if midSum < target {
			combine(candidates[i+1:], append(mid, candidates[i]), midSum, target, result)
		} else {
			return
		}
		midSum -= candidates[i]
	}
}
