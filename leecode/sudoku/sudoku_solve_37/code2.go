// Time        : 2019/07/01
// Description :

package sudoku_solve_37

func solveSudoku(board [][]byte) {
	solver(board, 0, 0)
}

var entries = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func solver(board [][]byte, i, j int) bool {
	if i == 9 {
		return true
	}
	ni, nj := next(i, j)
	if board[i][j] != '.' {
		return solver(board, ni, nj)
	}
	for _, b := range entries {
		if subBox(board, b, i/3*3, j/3*3) && line(board, b, i) && col(board, b, j) {
			board[i][j] = b
			if solver(board, ni, nj) {
				return true
			}
			board[i][j] = '.'
		}
	}
	return false
}

func subBox(board [][]byte, b byte, x, y int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+x][j+y] == b {
				return false
			}
		}
	}
	return true
}

func line(board [][]byte, b byte, i int) bool {
	for j := 0; j < 9; j++ {
		if board[i][j] == b {
			return false
		}
	}
	return true
}

func col(board [][]byte, b byte, j int) bool {
	for i := 0; i < 9; i++ {
		if board[i][j] == b {
			return false
		}
	}
	return true
}

func next(i, j int) (int, int) {
	if j < 8 {
		return i, j + 1
	}
	return i + 1, 0
}
