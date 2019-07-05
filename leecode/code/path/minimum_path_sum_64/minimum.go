// Time        : 2019/07/04
// Description :

package minimum_path_sum_64

// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.
//
// Note: You can only move either down or right at any point in time.
//
// Example:
//
// Input:
// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
// Output: 7
// Explanation: Because the path 1→3→1→1→1 minimizes the sum.

func minPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			if i == row-1 && j == col-1 {
				continue
			}
			if i == row-1 {
				grid[i][j] = grid[i][j+1] + grid[i][j]
			} else if j == col-1 {
				grid[i][j] = grid[i+1][j] + grid[i][j]
			} else {
				grid[i][j] = min(grid[i+1][j], grid[i][j+1]) + grid[i][j]
			}
		}
	}
	return grid[0][0]
}
