// Time        : 2019/07/25
// Description :

package surrounded_regions_130

// Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.
//
// A region is captured by flipping all 'O's into 'X's in that surrounded region.
//
// Example:
//
// X X X X
// X O O X
// X X O X
// X O X X
// After running your function, the board should be:
//
// X X X X
// X X X X
// X X X X
// X O X X
// Explanation:
//
// Surrounded regions shouldn’t be on the border,
// which means that any 'O' on the border of the board are not flipped to 'X'.
// Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'.
// Two cells are connected if they are adjacent cells connected horizontally or vertically.

var di = []int{0, 1, 0, -1}
var dj = []int{1, 0, -1, 0}

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	i, j, dx := 0, -1, 0
	for dx < 4 {
		ni := i + di[dx]
		nj := j + dj[dx]
		if ni >= len(board) || ni < 0 || nj >= len(board[0]) || nj < 0 {
			dx++
			continue
		}
		i, j = ni, nj
		if board[i][j] == 'O' {
			spring(board, i, j)
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}

func spring(board [][]byte, i, j int) {
	board[i][j] = 'A'
	for dx := 0; dx < 4; dx++ {
		ni, nj := i+di[dx], j+dj[dx]
		if ni >= len(board) || ni < 0 || nj >= len(board[0]) || nj < 0 {
			continue
		}
		if board[ni][nj] == 'O' {
			spring(board, ni, nj)
		}
	}
}
