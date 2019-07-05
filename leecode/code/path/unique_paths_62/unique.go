// Time        : 2019/07/04
// Description :

package unique_paths_62

//A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
//
//The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
//
//How many possible unique paths are there?
//
//Above is a 7 x 3 grid. How many possible unique paths are there?
//
//Note: m and n will be at most 100.
//
//Example 1:
//
//Input: m = 3, n = 2
//Output: 3
//Explanation:
//From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
//1. Right -> Right -> Down
//2. Right -> Down -> Right
//3. Down -> Right -> Right
//Example 2:
//
//Input: m = 7, n = 3
//Output: 28

func stupidPaths(m int, n int) int {
	var result int
	walk(1, 1, m, n, &result)
	return result
}

var dx = []int{1, 0}
var dy = []int{0, 1}

func walk(x, y, m, n int, result *int) {
	if x == m && y == n {
		*result++
		return
	}
	for i := 0; i < 2; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx > m || ny > n {
			continue
		}
		walk(nx, ny, m, n, result)
	}
}

func uniquePaths(m int, n int) int {
	paths := make([][]int, n+1)
	for i := range paths {
		paths[i] = make([]int, m+1)
	}
	paths[n-1][m-1] = 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i == n-1 && j == m-1 {
				continue
			}
			paths[i][j] = paths[i+1][j] + paths[i][j+1]
		}
	}
	return paths[0][0]
}
