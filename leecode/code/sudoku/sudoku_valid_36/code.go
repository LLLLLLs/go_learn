// Time        : 2019/07/01
// Description :

package sudoku_valid_36

// Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//
// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.
//
// A partially filled sudoku which is valid.
//
// The Sudoku board could be partially filled, where empty cells are filled with the character '.'.
// Example 1:
// Input:
// [
//   ["5","3",".",".","7",".",".",".","."],
//   ["6",".",".","1","9","5",".",".","."],
//   [".","9","8",".",".",".",".","6","."],
//   ["8",".",".",".","6",".",".",".","3"],
//   ["4",".",".","8",".","3",".",".","1"],
//   ["7",".",".",".","2",".",".",".","6"],
//   [".","6",".",".",".",".","2","8","."],
//   [".",".",".","4","1","9",".",".","5"],
//   [".",".",".",".","8",".",".","7","9"]
// ]
// Output: true

const empty = '.'

func isValidSudoku(board [][]byte) bool {
	judge := func(m map[byte]bool, key byte) bool {
		if key == empty {
			return true
		}
		if m[key] {
			return false
		}
		m[key] = true
		return true
	}
	subBox := func(x, y int) bool {
		var m = make(map[byte]bool)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if !judge(m, board[i+x][j+y]) {
					return false
				}
			}
		}
		return true
	}
	lineAndCol := func(i int) bool {
		var line = make(map[byte]bool)
		var col = make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if !judge(line, board[i][j]) {
				return false
			}
			if !judge(col, board[j][i]) {
				return false
			}
		}
		return true
	}
	for i := 0; i < 9; i++ {
		if !subBox(i%3*3, i/3*3) || !lineAndCol(i) {
			return false
		}
	}
	return true
}
