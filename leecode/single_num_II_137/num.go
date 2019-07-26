// Time        : 2019/07/26
// Description :

package single_num_II_137

// Given a non-empty array of integers, every element appears three times except for one,
// which appears exactly once. Find that single one.
//
// Note:
//
// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
//
// Example 1:
//
// Input: [2,2,3,2]
// Output: 3
// Example 2:
//
// Input: [0,1,0,1,0,1,99]
// Output: 99

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for i := range nums {
		ones = (ones ^ nums[i]) & (^twos)
		twos = (twos ^ nums[i]) & (^ones)
	}
	return ones
}

// 原文链接：https://leetcode.com/problems/single-number-ii/discuss/43295/Detailed-explanation-and-generalization-of-the-bitwise-operation-method-for-single-numbers
// 思路如下：就两个字：牛逼
// The code seems tricky and hard to understand at first glance.
// However, if you consider the problem in Boolean algebra form, everything becomes clear.
//
// What we need to do is to store the number of '1's of every bit.
// Since each of the 32 bits follow the same rules, we just need to consider 1 bit.
// We know a number appears 3 times at most, so we need 2 bits to store that.
// Now we have 4 state, 00, 01, 10 and 11, but we only need 3 of them.
//
// In this solution, 00, 01 and 10 are chosen.
// Let 'ones' represents the first bit, 'twos' represents the second bit.
// Then we need to set rules for 'ones' and 'twos' so that they act as we hopes.
// The complete loop is 00->10->01->00(0->1->2->3/0).
//
// For 'ones', we can get 'ones = ones ^ A[i]; if (twos == 1) then ones = 0',
// that can be tansformed to 'ones = (ones ^ A[i]) & ~twos'.
//
// Similarly, for 'twos', we can get 'twos = twos ^ A[i]; if (ones* == 1) then twos = 0' and 'twos = (twos ^ A[i]) & ~ones'.
// Notice that 'ones*' is the value of 'ones' after calculation, that is why twos is calculated later.
