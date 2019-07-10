// Time        : 2019/07/09
// Description :

package climbingstairs_70

// You are climbing a stair case. It takes n steps to reach to the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Note: Given n will be a positive integer.
//
// Example 1:
//
// Input: 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps
// Example 2:
//
// Input: 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step

func climbStairs(n int) int {
	fn := [2]int{1, 2}
	if n < 3 {
		return fn[n-1]
	}
	for i := 2; i < n; i++ {
		fn[0], fn[1] = fn[1], fn[0]+fn[1]
	}
	return fn[1]
}
