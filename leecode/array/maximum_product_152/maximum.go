// Time        : 2019/07/31
// Description :

package maximum_product_152

import "math"

// Given an integer array nums,
// find the contiguous subarray within an array (containing at least one number)
// which has the largest product.
//
// Example 1:
//
// Input: [2,3,-2,4]
// Output: 6
// Explanation: [2,3] has the largest product 6.
// Example 2:
//
// Input: [-2,0,-1]
// Output: 0
// Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

// 首先该子序列一定不包含0（除非nums是[负，0，负],或全为0的情况）
// 对于一个不包含0的序列，子序列下标一定从0开始或终止于n:
// 	假设数组m[i,j],最大乘积子序列max[m,n],其中,m>i&&n<j,对于m-1和n+1有以下情况:
//  	m-1与n+1同号:则max[m,n]可扩展至max[m-1,n+1];
//		m-1与n+1异号:则max[m,n]可扩展至max[m-1,n]或max[m,n+1].
//	上述操作直到遇到数组m边界i,j为止。
// 因此算法可以计算从左到右与从右到左的最大值，取2者最大值。
func maxProduct(nums []int) int {
	left := maxLeft(nums)
	reverse(nums)
	right := maxLeft(nums)
	if left > right {
		return left
	}
	return right
}

func maxLeft(nums []int) int {
	res := math.MinInt32
	product := 1
	for i := range nums {
		product *= nums[i]
		if product > res {
			res = product
		}
		if product == 0 {
			product = 1
		}
	}
	return res
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-1-i], nums[i]
	}
}
