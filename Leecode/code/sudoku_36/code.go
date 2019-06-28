// Time        : 2019/06/28
// Description :

package sudoku_36

import (
	"encoding/json"
	"go_learn/utils"
)

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

type pos struct {
	i int
	j int
}

func isValidSudoku(board [][]byte) bool {
	var sudokuMap = make(map[pos]map[byte]bool)
	generateNode := func(i, j int) map[byte]bool {
		node := make(map[byte]bool)
		for k := byte(1); k <= 9; k++ {
			node['0'+k] = true
		}
		for i := 0; i < 9; i++ {
			delete(node, board[i][j])
		}
		for j := 0; j < 9; j++ {
			delete(node, board[i][j])
		}
		return node
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				node := generateNode(i, j)
				sudokuMap[pos{i: i, j: j}] = node
			}
		}
	}
	node, ok := getNode(sudokuMap, 0, 0)
	if !ok {
		return solve(marshal(sudokuMap), '0', 0, 0)
	}
	for k := range node {
		if solve(marshal(sudokuMap), k, 0, 0) {
			return true
		}
	}
	return false
}

func solve(data []byte, key byte, i, j int) bool {
	if i == 10 {
		return true
	}
	m := unmarshal(data)
	node, ok := getNode(m, i, j)
	if !ok {
		i, j = next(i, j)
		return solve(data, '0', i, j)
	}
	for k := range node {

	}
	return false
}

func deleteKey(m map[pos]map[byte]bool, key byte, i, j int) bool {
	if key == '0' {
		return true
	}
	for k := 0; k < 9; k++ {
		node, ok := getNode(m, k, j)
		if ok {
			delete(node, key)
		}
	}
	for k := 0; k < 9; k++ {
		node, ok := getNode(m, i, k)
		if ok {
			delete(node, key)
		}
	}
}

func next(i, j int) (int, int) {
	if j < 8 {
		return i, j + 1
	}
	return i + 1, 0
}

func getNode(m map[pos]map[byte]bool, i, j int) (node map[byte]bool, ok bool) {
	node, ok = m[pos{i: i, j: j}]
	return
}

func marshal(m map[pos]map[byte]bool) []byte {
	data, err := json.Marshal(m)
	utils.OkOrPanic(err)
	return data
}

func unmarshal(data []byte) map[pos]map[byte]bool {
	var m map[pos]map[byte]bool
	err := json.Unmarshal(data, &m)
	utils.OkOrPanic(err)
	return m
}
