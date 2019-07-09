// Time        : 2019/07/04
// Description :

package unique_paths_II_63

// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
//
// The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
//
// Now consider if some obstacles are added to the grids. How many unique paths would there be?
//
// An obstacle and empty space is marked as 1 and 0 respectively in the grid.
//
// Note: m and n will be at most 100.
//
// Example 1:
//
// Input:
// [
//   [0,0,0],
//   [0,1,0],
//   [0,0,0]
// ]
// Output: 2
// Explanation:
// There is one obstacle in the middle of the 3x3 grid above.
// There are two ways to reach the bottom-right corner:
// 1. Right -> Right -> Down -> Down
// 2. Down -> Down -> Right -> Right

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	paths := make([][]int, row+1)
	for i := range paths {
		paths[i] = make([]int, col+1)
	}
	if obstacleGrid[row-1][col-1] == 1 {
		return 0
	}
	paths[row-1][col-1] = 1
	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			if i == row-1 && j == col-1 {
				continue
			}
			if obstacleGrid[i][j] != 1 {
				paths[i][j] = paths[i+1][j] + paths[i][j+1]
			}
		}
	}
	return paths[0][0]
}
