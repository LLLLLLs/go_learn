// Time        : 2019/06/28
// Description :

package sudoku_solve_37

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

const empty = '.'

var sortedPoint [10][][2]int
var isFilled func(i, j int) bool

func initSortedPoint() {
	for i := range sortedPoint {
		sortedPoint[i] = make([][2]int, 0)
	}
}

func nextPoint(x, y int) (int, int, int, int) {
	for x != 10 && y >= len(sortedPoint[x]) {
		x++
		y = 0
	}
	if x == 10 {
		return x, y, -1, -1
	}
	y++
	return x, y, sortedPoint[x][y-1][0], sortedPoint[x][y-1][1]
}

func isValidSudoku(board [][]byte) {
	initSortedPoint()
	isFilled = func(i, j int) bool {
		return board[i][j] != empty
	}
	var sudokuMap [9][9]map[byte]bool
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
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				delete(node, board[i/3*3+x][j/3*3+y])
			}
		}
		return node
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				node := generateNode(i, j)
				sudokuMap[i][j] = node
				sortedPoint[len(node)] = append(sortedPoint[len(node)], [2]int{i, j})
			}
		}
	}
	x, y, i, j := nextPoint(0, 0)
	node := getNode(sudokuMap, i, j)
	for k := range node {
		board[i][j] = k
		if solve(board, marshal(sudokuMap), k, x, y, i, j) {
			return
		}
		board[i][j] = empty
	}
	return
}

func solve(board [][]byte, data []byte, key byte, x, y, i, j int) bool {
	m := unmarshal(data)
	if !deleteKey(m, key, i, j) {
		return false
	}
	data = marshal(m)
	x, y, i, j = nextPoint(x, y)
	if i == -1 {
		//printSudoku(board)
		return true
	}
	node := getNode(m, i, j)
	for k := range node {
		board[i][j] = k
		if solve(board, data, k, x, y, i, j) {
			return true
		}
		board[i][j] = empty
	}
	return false
}

func deleteKey(m [9][9]map[byte]bool, key byte, i, j int) bool {
	if key == '0' {
		return true
	}
	deleteNode := func(x, y int) bool {
		if isFilled(x, y) {
			return true
		}
		node := getNode(m, x, y)
		if len(node) == 0 {
			return false
		}
		delete(node, key)
		return true
	}
	// 横竖
	for k := 0; k < 9; k++ {
		if !deleteNode(k, j) {
			return false
		}
		if !deleteNode(i, k) {
			return false
		}
	}
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if !deleteNode(i/3*3+x, j/3*3+y) {
				return false
			}
		}
	}
	return true
}

func getNode(m [9][9]map[byte]bool, i, j int) map[byte]bool {
	return m[i][j]
}

func marshal(m [9][9]map[byte]bool) []byte {
	io := bytes.Buffer{}
	enc := gob.NewEncoder(&io)
	err := enc.Encode(m)
	if err != nil {
		panic(err)
	}
	return io.Bytes()
}

func unmarshal(data []byte) [9][9]map[byte]bool {
	var m [9][9]map[byte]bool
	io := bytes.NewBuffer(data)
	dec := gob.NewDecoder(io)
	err := dec.Decode(&m)
	if err != nil {
		panic(err)
	}
	return m
}

func printSudoku(board [][]byte) {
	for _, list := range board {
		for i := range list {
			list[i] -= '0'
		}
	}
	for _, list := range board {
		fmt.Println(list)
	}
}
