// Time        : 2019/07/10
// Description :

package word_search_79

import (
	"golearn/util"
)

// Given a 2D board and a word, find if the word exists in the grid.
//
// The word can be constructed from letters of sequentially adjacent cell,
// where "adjacent" cells are those horizontally or vertically neighboring.
// The same letter cell may not be used more than once.
//
// Example:
//
// board =
// [
//   ['A','B','C','E'],
//   ['S','F','C','S'],
//   ['A','D','E','E']
// ]
//
// Given word = "ABCCED", return true.
// Given word = "SEE", return true.
// Given word = "ABCB", return false.

func exist(board [][]byte, word string) bool {
	path := make([][]byte, len(board))
	for i := range path {
		path[i] = make([]byte, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != word[0] {
				continue
			}
			path[i][j] = 1
			if backtrack(board, path, i, j, word[1:]) {
				util.Print2DimensionList(path)
				return true
			}
			path[i][j] = 0
		}
	}
	return false
}

var di = []int{0, 1, 0, -1}
var dj = []int{1, 0, -1, 0}

func backtrack(board [][]byte, path [][]byte, i, j int, word string) bool {
	if len(word) == 0 {
		return true
	}
	for k := 0; k < 4; k++ {
		ni, nj := i+di[k], j+dj[k]
		if ni >= 0 && nj >= 0 && ni < len(board) && nj < len(board[ni]) && path[ni][nj] == 0 && board[ni][nj] == word[0] {
			path[ni][nj] = path[i][j] + 1
			if backtrack(board, path, ni, nj, word[1:]) {
				return true
			}
			path[ni][nj] = 0
		}
	}
	return false
}
