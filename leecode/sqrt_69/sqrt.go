// Time        : 2019/07/09
// Description :

package sqrt_69

// Implement int sqrt(int x).
//
// Compute and return the square root of x, where x is guaranteed to be a non-negative integer.
//
// Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.
//
// Example 1:
//
// Input: 4
// Output: 2
// Example 2:
//
// Input: 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since
//              the decimal part is truncated, 2 is returned.

func mySqrt(x int) int {
	left, right := 1, 1
	for left*left < x {
		left *= 2
	}
	if left*left == x {
		return left
	}
	left, right = left/2, left
	for right-left != 1 {
		next := left + (right-left)/2
		if next*next == x {
			return next
		} else if next*next > x {
			right = next
		} else {
			left = next
		}
	}
	return left
}
