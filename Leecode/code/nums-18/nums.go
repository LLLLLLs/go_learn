// Time        : 2019/01/21
// Description :

package nums_18

import "sort"

// Given an array nums of n integers and an integer target,
// are there elements a, b, c, and d in nums such that a + b + c + d = target?
// Find all unique quadruplets in the array which gives the sum of target.

// 暴力遍历 -- 四重循环
func fourSum1(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	for a := 0; a < len(nums)-3; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < len(nums)-2; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			for c := b + 1; c < len(nums)-1; c++ {
				if c > b+1 && nums[c] == nums[c-1] {
					continue
				}
				for d := c + 1; d < len(nums); d++ {
					if d > c+1 && nums[d] == nums[d-1] {
						continue
					}
					if nums[a]+nums[b]+nums[c]+nums[d] == target {
						result = append(result, []int{nums[a], nums[b], nums[c], nums[d]})
					}
				}
			}
		}
	}
	return result
}

// 分2次2重循环
func fourSum2(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	length := len(nums)
	// 第一次二重循环 获取两数之和
	twoSum := make(map[int][][2]int)
	for i := 0; i < length-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a := nums[i]
		for j := i + 1; j < length; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			b := nums[j]
			if l, ok := twoSum[a+b]; !ok {
				l = [][2]int{{i, j}}
				twoSum[a+b] = l
			} else {
				l = append(l, [2]int{i, j})
				twoSum[a+b] = l
			}
		}
	}

	for i := 0; i < length-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a := nums[i]
		for j := i + 1; j < length; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			b := nums[j]
			list, ok := twoSum[target-a-b]
			if !ok {
				continue
			}
			for _, l := range list {
				if l[0] == i || l[0] == j || l[1] == i || l[1] == j {
					continue
				}
				result = append(result, []int{nums[l[0]], nums[l[1]], a, b})
			}
		}
	}
	for _, r := range result {
		sort.Ints(r)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0] ||
			result[i][0] == result[j][0] && result[i][1] < result[j][1] ||
			result[i][0] == result[j][0] && result[i][1] == result[j][1] && result[i][2] < result[j][2] ||
			result[i][0] == result[j][0] && result[i][1] == result[j][1] && result[i][2] == result[j][2] && result[i][3] < result[j][3]
	})
	rtn := make([][]int, 0)
	index := 0
	for _, r := range result {
		if len(rtn) == 0 {
			rtn = append(rtn, r)
			continue
		}
		repeat := true
		for i := range r {
			if r[i] != rtn[index][i] {
				repeat = false
				break
			}
		}
		if !repeat {
			rtn = append(rtn, r)
			index++
		}
	}
	return rtn
}

// use two pointers to find the fourth number
func fourSum3(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	ell := len(nums)

	sort.Ints(nums)
	for i := 0; i < ell-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < ell-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for k, l := j+1, ell-1; k < l; {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				} else if sum < target {
					k++
				} else {
					l--
				}
			}
		}
	}

	return ans
}

func fourSum4(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	length := len(nums)
	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k, l := j+1, length-1; k < l; {
				if nums[i]+nums[j]+nums[k]+nums[l] < target {
					k++
				} else if nums[i]+nums[j]+nums[k]+nums[l] > target {
					l--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				}
			}
		}
	}
	return result
}
