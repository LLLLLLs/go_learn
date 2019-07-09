// Time        : 2019/07/03
// Description :

package rotate_image_48

import (
	"go_learn/utils"
	"testing"
)

func TestRotate(t *testing.T) {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	utils.Print2DimensionList(matrix)
	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	utils.Print2DimensionList(matrix)
}
