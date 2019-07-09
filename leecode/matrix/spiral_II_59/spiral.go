// Time        : 2019/07/04
// Description :

package spiral_II_59

// Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.
//
// Example:
//
// Input: 3
// Output:
// [
//  [ 1, 2, 3 ],
//  [ 8, 9, 4 ],
//  [ 7, 6, 5 ]
// ]

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	var dx = []int{0, 1, 0, -1}
	var dy = []int{1, 0, -1, 0}
	var x, y, di = 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[x][y] = i
		nx, ny := x+dx[di], y+dy[di]
		if nx == n || nx == -1 || ny == n || ny == -1 || matrix[nx][ny] != 0 {
			di++
			di = di % 4
			nx, ny = x+dx[di], y+dy[di]
		}
		x, y = nx, ny
	}
	return matrix
}
