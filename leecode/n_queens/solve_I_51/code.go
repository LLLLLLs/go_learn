// Time        : 2019/07/03
// Description :

package solve_I_51

import (
	"strings"
)

// The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.
//
// Given an integer n, return all distinct solutions to the n-queens puzzle.
//
// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.
//
// Example:
//
// Input: 4
// Output: [
//  [".Q..",  // Solution 1
//   "...Q",
//   "Q...",
//   "..Q."],
//
//  ["..Q.",  // Solution 2
//   "Q...",
//   "...Q",
//   ".Q.."]
// ]
// Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	recursive([]string{}, n, &result)
	return result
}

func recursive(mid []string, n int, result *[][]string) {
	if len(mid) == n {
		*result = append(*result, mid)
		return
	}
	for j := 0; j < n; j++ {
		hasQueen := false
		for i := range mid {
			if mid[i][j] == 'Q' {
				hasQueen = true
				break
			}
		}
		for i := 0; !hasQueen && i < n; i++ {
			skewI := len(mid) - abs(i-j)
			if 0 <= skewI && skewI < len(mid) && mid[skewI][i] == 'Q' {
				hasQueen = true
				break
			}
		}
		if !hasQueen {
			recursive(append(append([]string{}, mid...), genCol(j, n)), n, result)
		}
	}
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func genCol(index, n int) string {
	builder := strings.Builder{}
	for i := 0; i < n; i++ {
		if i == index {
			builder.WriteByte('Q')
		} else {
			builder.WriteByte('.')
		}
	}
	return builder.String()
}
