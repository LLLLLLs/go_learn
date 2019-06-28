// Time        : 2019/06/26
// Description :

package divide_two_integer_29

import "math"

// Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.
//
// Return the quotient after dividing dividend by divisor.
//
// The integer division should truncate toward zero.
//
// Example 1:
//
// Input: dividend = 10, divisor = 3
// Output: 3

// Example 2:
//
// Input: dividend = 7, divisor = -3
// Output: -2

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend == 0 {
		return 0
	}
	var sign = -1
	if dividend^divisor >= 0 {
		sign = 1
	}
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	dvd, dvs := abs(dividend), abs(divisor)
	var result int
	for dvd >= dvs {
		tmp, m := dvs, 1
		for dvd > tmp<<1 {
			tmp, m = tmp<<1, m<<1
		}
		result += m
		dvd -= tmp
	}
	return result * sign
}
