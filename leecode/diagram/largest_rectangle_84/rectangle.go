// Time        : 2019/07/11
// Description :

package largest_rectangle_84

// Given n non-negative integers representing the histogram's bar height where the width of each bar is 1,
// find the area of largest rectangle in the histogram.
//
// Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].
//
// The largest rectangle is shown in the shaded area, which has area = 10 unit.
//
// Example:
//
// Input: [2,1,5,6,2,3]
// Output: 10

func largestRectangleArea(heights []int) int {
	heights = append(heights, -1)
	stack := make([]int, 0, len(heights))
	result := 0
	for i := 0; i < len(heights); {
		if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[top]
			w := i
			if len(stack) != 0 {
				w = i - stack[len(stack)-1] - 1
			}
			result = max(result, h*w)
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
