// Time        : 2019/07/23
// Description :

package III_123

import "math"

// Say you have an array for which the ith element is the price of a given stock on day i.
//
// Design an algorithm to find the maximum profit. You may complete at most two transactions.
//
// Note: You may not engage in multiple transactions at the same time
// (i.e., you must sell the stock before you buy again).
//
// Example 1:
//
// Input: [3,3,5,0,0,3,1,4]
// Output: 6
// Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
//              Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
// Example 2:
//
// Input: [1,2,3,4,5]
// Output: 4
// Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
//              Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
//              engaging multiple transactions at the same time. You must sell before buying again.
// Example 3:
//
// Input: [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transaction is done, i.e. max profit = 0.

func maxProfit(prices []int) int {
	max := 0
	j := -1
	first := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] && j == -1 {
			continue
		}
		if j == -1 {
			j = i - 1
			continue
		}
		if prices[i] > prices[i-1] {
			continue
		}
		if prices[i-1]-prices[j] > first {
			first = prices[i-1] - prices[j]
			profit := first + _maxProfit(prices[i:])
			if profit > max {
				max = profit
			}
		}
		for i < len(prices) && prices[i] < prices[i-1] {
			if prices[i] < prices[j] {
				j = i
			}
			i++
		}
	}
	if max == 0 && len(prices) > 0 && j != -1 && prices[len(prices)-1] > prices[j] {
		max = prices[len(prices)-1] - prices[j]
	}
	return max
}

func _maxProfit(prices []int) int {
	maxPrice, profit := 0, 0
	for i := len(prices) - 1; i >= 0; i-- {
		if maxPrice-prices[i] > profit {
			profit = maxPrice - prices[i]
		}
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		}
	}
	return profit
}

// 看得懂想不到
func maxProfit2(prices []int) int {
	firstbuy, secondbuy, firstsell, secondsell := math.MinInt32, math.MinInt32, 0, 0
	for i := 0; i < len(prices); i++ {
		firstbuy = max(firstbuy, -prices[i])
		firstsell = max(firstsell, prices[i]+firstbuy)
		secondbuy = max(secondbuy, firstsell-prices[i])
		secondsell = max(secondsell, prices[i]+secondbuy)
	}
	return secondsell
}

func max(num1 int, num2 int) int {
	if num1 < num2 {
		return num2
	} else {
		return num1
	}
}
