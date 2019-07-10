// Time        : 2019/07/10
// Description :

package set_zeroes_73

// Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.
//
// Example 1:
//
// Input:
// [
//   [1,1,1],
//   [1,0,1],
//   [1,1,1]
// ]
// Output:
// [
//   [1,0,1],
//   [0,0,0],
//   [1,0,1]
// ]
// Example 2:
//
// Input:
// [
//   [0,1,2,0],
//   [3,4,5,2],
//   [1,3,1,5]
// ]
// Output:
// [
//   [0,0,0,0],
//   [0,4,5,0],
//   [0,3,1,0]
// ]

func setZeroes(matrix [][]int) {
	var flag = make([][]bool, len(matrix))
	for i := range flag {
		flag[i] = make([]bool, len(matrix[0]))
	}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				for k := 0; k < len(matrix[0]); k++ {
					flag[i][k] = true
				}
				for k := 0; k < len(matrix); k++ {
					flag[k][j] = true
				}
			}
		}
	}
	for i := range flag {
		for j := range flag[i] {
			if flag[i][j] {
				matrix[i][j] = 0
			}
		}
	}
}

func setZeroesMap(matrix [][]int) {
	var row = make(map[int]struct{})
	var col = make(map[int]struct{})
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				row[i] = struct{}{}
				col[j] = struct{}{}
			}
		}
	}
	for r := range row {
		for c := 0; c < len(matrix[0]); c++ {
			matrix[r][c] = 0
		}
	}
	for c := range col {
		for r := 0; r < len(matrix); r++ {
			matrix[r][c] = 0
		}
	}
}
