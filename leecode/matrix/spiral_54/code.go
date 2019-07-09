// Time        : 2019/07/04
// Description :

package spiral_54

// Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
//
// Example 1:
//
// Input:
// [
//  [ 1, 2, 3 ],
//  [ 4, 5, 6 ],
//  [ 7, 8, 9 ]
// ]
// Output: [1,2,3,6,9,8,7,4,5]
// Example 2:
//
// Input:
// [
//   [1, 2, 3, 4],
//   [5, 6, 7, 8],
//   [9,10,11,12]
// ]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]

func spiralOrder(matrix [][]int) []int {
	var result = make([]int, 0)
	if len(matrix) == 0 {
		return result
	}
	var dx = []int{0, 1, 0, -1}
	var dy = []int{1, 0, -1, 0}
	var seen = make([][]bool, len(matrix))
	for i := range matrix {
		seen[i] = make([]bool, len(matrix[i]))
	}
	var x, y, di = 0, 0, 0
	for i := 0; i < len(matrix)*len(matrix[0]); i++ {
		result = append(result, matrix[x][y])
		nx, ny := x+dx[di], y+dy[di]
		if nx >= len(matrix) || nx < 0 || ny >= len(matrix[0]) || ny < 0 || seen[nx][ny] {
			di++
			di = di % 4
			nx, ny = x+dx[di], y+dy[di]
		}
		seen[x][y] = true
		x, y = nx, ny
	}
	return result
}
