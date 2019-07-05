// Time        : 2019/07/03
// Description :

package rotate_image_48

// You are given an n x n 2D matrix representing an image.
//
// Rotate the image by 90 degrees (clockwise).
//
// Note:
//
// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
//
// Example 1:
//
// Given input matrix =
// [
//   [1,2,3],
//   [4,5,6],
//   [7,8,9]
// ],
//
// rotate the input matrix in-place such that it becomes:
// [
//   [7,4,1],
//   [8,5,2],
//   [9,6,3]
// ]
// Example 2:
//
// Given input matrix =
// [
//   [ 5, 1, 9,11],
//   [ 2, 4, 8,10],
//   [13, 3, 6, 7],
//   [15,14,12,16]
// ],
//
// rotate the input matrix in-place such that it becomes:
// [
//   [15,13, 2, 5],
//   [14, 3, 4, 1],
//   [12, 6, 8, 9],
//   [16, 7,10,11]
// ]

func rotate(matrix [][]int) {
	width := len(matrix)
	next := func(x, y int) (int, int) {
		return y, width - x - 1
	}
	for i := 0; i < width/2; i++ {
		for j := i; j < (width-i)-1; j++ {
			x, y := i, j
			var tmp = matrix[x][y]
			for k := 0; k < 4; k++ {
				nx, ny := next(x, y)
				matrix[nx][ny], tmp = tmp, matrix[nx][ny]
				x, y = nx, ny
			}
		}
	}
}
