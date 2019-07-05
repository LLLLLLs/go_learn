/*
Author      : lls
Time        : 2018/10/30
Description :
*/

package code

import "fmt"

//编写一个程序，通过已填充的空格来解决数独问题。
//
//一个数独的解法需遵循如下规则：
//
//数字 1-9 在每一行只能出现一次。
//数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
//空白格用 '.' 表示。
//
//
//
//一个数独。
//
//
//
//答案被标成红色。

type node []int

var sudokuMay [9][9]node

func SolveSudoku(Sudoku [9][9]int) {
	n := inited(Sudoku)
	SudokuSure, _ := sure(sudokuMay)
	for n > 0 {
		n = Subinit(SudokuSure)
		SudokuSure, _ = sure(sudokuMay)
	}
	Output()
	fmt.Println(isEnable(sudokuMay))
}
func isEnable(tn [9][9]node) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if len(tn[i][j]) == 0 {
				return false
			}
		}
	}
	return true
}
func sure(may [9][9]node) (sure [9][9]int, n int) {
	n = 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if len(may[i][j]) == 1 {
				sure[i][j] = may[i][j][0]
				n++
			} else {
				sure[i][j] = 0
			}
		}
	}
	return
}

func inited(Sud [9][9]int) (changeCount int) {
	tmp := 0
	changeCount = 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if Sud[i][j] != 0 {
				sudokuMay[i][j] = append(sudokuMay[i][j], Sud[i][j])
			} else {
				for k := 0; k < 9; k++ {
					sudokuMay[i][j] = append(sudokuMay[i][j], k+1)
				}
				sudokuMay[i][j], tmp = excludeMay(i, j, sudokuMay[i][j], Sud)
				changeCount += tmp
			}
		}
	}
	return
}
func Subinit(Sud [9][9]int) (changeCount int) {
	tmp := 0
	changeCount = 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if Sud[i][j] != 0 {
				sudokuMay[i][j][0] = Sud[i][j]
			} else {
				sudokuMay[i][j], tmp = excludeMay(i, j, sudokuMay[i][j], Sud)
				changeCount += tmp
			}
		}
	}
	return
}
func excludeMay(ti, tj int, t node, S [9][9]int) (rmay node, changeCount int) {
	changeCount = 0
	var tmpChangeCount int
	for i := 0; i < 9; i++ {
		if S[i][tj] != 0 {
			t, tmpChangeCount = exclude(t, S[i][tj])
			changeCount += tmpChangeCount
		}
		if S[ti][i] != 0 {
			t, tmpChangeCount = exclude(t, S[ti][i])
			changeCount += tmpChangeCount
		}
	}
	for k := (ti / 3) * 3; k < (ti/3)*3+3; k++ {
		for l := (tj / 3) * 3; l < (tj/3)*3+3; l++ {
			if S[k][l] != 0 {
				t, tmpChangeCount = exclude(t, S[k][l])
				changeCount += tmpChangeCount
			}
		}
	}
	rmay = t
	return
}
func excludeFirstOne(smay node, n int) (rmay node, changeCount int) {
	changeCount = 0
	rmay = smay
	for i := 0; i < len(smay); i++ {
		if smay[i] == n {
			changeCount++
			rmay = append(smay[:i], smay[i+1:]...)
			return
		}
		if i == len(smay)-1 {
			return
		}
	}
	return
}
func exclude(smay node, n int) (tmp node, changeCount int) {
	var nc int
	changeCount = 0
	tmp, nc = excludeFirstOne(smay, n)
	for nc > 0 {
		tmp, nc = excludeFirstOne(tmp, n)
		changeCount++
	}
	return
}
func Output() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(sudokuMay[i][j])
		}
		fmt.Println("")
	}
}
