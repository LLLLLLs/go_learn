// Time        : 2019/11/08
// Description :

package conwaylife

const MAX_BOARD_SIZE = 64

var di = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dj = []int{-1, 0, 1, -1, 1, -1, 0, 1}

var emptyBoard = make([][]bool, 0)
var endBoard = [][]bool{
	{false, false},
	{false, false},
}

func NextStep(now [][]bool) [][]bool {
	if len(now) == 0 {
		return nil
	}
	var count = make([][]byte, len(now))
	for i := range count {
		count[i] = make([]byte, len(now[i]))
		for j := range count[i] {
			for k := range di {
				aliveCount(now, i+di[k], j+dj[k], &count[i][j])
			}
		}
	}
	next := nextBoard(now, count)
	next = formatBoard(next)
	if len(next) == 0 || equal(now, next) {
		return emptyBoard
	}
	return next
}

// 收缩或扩张领地
func formatBoard(board [][]bool) [][]bool {
	if len(board) == 0 {
		return emptyBoard
	}
	board = topTrim(board)
	board = bottomTrim(board)
	board = leftTrim(board)
	board = rightTrim(board)
	if equal(board, endBoard) {
		return emptyBoard
	}
	return board
}

func topTrim(board [][]bool) [][]bool {
	var row int
	for row = range board {
		for col := range board[row] {
			if board[row][col] {
				goto outLoop
			}
		}
	}
outLoop:
	if row == 0 && len(board) < MAX_BOARD_SIZE {
		newLine := make([]bool, len(board[0]))
		board = append([][]bool{newLine}, board...)
	}
	if row > 1 {
		board = board[row-1:]
	}
	return board
}

func bottomTrim(board [][]bool) [][]bool {
	var row int
	for i := range board {
		row = len(board) - 1 - i
		for col := range board[row] {
			if board[row][col] {
				goto outLoop
			}
		}
	}
outLoop:
	if row == len(board)-1 && len(board) < MAX_BOARD_SIZE {
		newLine := make([]bool, len(board[0]))
		board = append(board, newLine)
	}
	if row < len(board)-2 {
		board = board[:row+2]
	}
	return board
}

func leftTrim(board [][]bool) [][]bool {
	var col int
	for col = range board[0] {
		for row := range board {
			if board[row][col] {
				goto outLoop
			}
		}
	}
outLoop:
	if col == 0 && len(board[0]) < MAX_BOARD_SIZE {
		for i := range board {
			board[i] = append([]bool{false}, board[i]...)
		}
	}
	if col > 1 {
		for i := range board {
			board[i] = board[i][col-1:]
		}
	}
	return board
}

func rightTrim(board [][]bool) [][]bool {
	var col int
	for i := range board[0] {
		col = len(board[0]) - 1 - i
		for row := range board {
			if board[row][col] {
				goto outLoop
			}
		}
	}
outLoop:
	if col == len(board[0])-1 && len(board[0]) < MAX_BOARD_SIZE {
		for i := range board {
			board[i] = append(board[i], false)
		}
	}
	if col < len(board[0])-2 {
		for i := range board {
			board[i] = board[i][:col+2]
		}
	}
	return board
}

func aliveCount(now [][]bool, i, j int, count *byte) {
	if i < 0 || j < 0 || i == len(now) || j == len(now[i]) {
		return
	}
	if now[i][j] {
		*count++
	}
}

func nextBoard(now [][]bool, count [][]byte) [][]bool {
	var next = make([][]bool, len(now))
	for i := range next {
		next[i] = make([]bool, len(now[i]))
		for j := range next[i] {
			if count[i][j] < 2 || count[i][j] > 3 {
				next[i][j] = false
			}
			if count[i][j] == 2 {
				next[i][j] = now[i][j]
			}
			if count[i][j] == 3 {
				next[i][j] = true
			}
		}
	}
	return next
}

func equal(a, b [][]bool) bool {
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return false
	}
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
