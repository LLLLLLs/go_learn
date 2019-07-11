// Time        : 2019/07/11
// Description :

package maximal_rectangle_85

// Given a 2D binary matrix filled with 0's and 1's,
// find the largest rectangle containing only 1's and return its area.
//
// Example:
//
// Input:
// [
//   ["1","0","1","0","0"],
//   ["1","0","1","1","1"],
//   ["1","1","1","1","1"],
//   ["1","0","0","1","0"]
// ]
// Output: 6

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	height := make([]int, len(matrix[0]))
	result := 0
	for i := range matrix {
		for j := range matrix[i] {
			if num := int(matrix[i][j] - '0'); num == 0 {
				height[j] = 0
			} else {
				height[j]++
			}
		}
		area := maxArea(height)
		result = max(area, result)
	}
	return result
}

func maxArea(heights []int) int {
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
