// Time        : 2019/07/04
// Description :

package spiral_54

import (
	"fmt"
	"testing"
)

func TestSpiral(t *testing.T) {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	result := spiralOrder(matrix)
	fmt.Println(result)
	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	result = spiralOrder(matrix)
	fmt.Println(result)
	matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
	}
	result = spiralOrder(matrix)
	fmt.Println(result)
	matrix = [][]int{
		{5},
		{2},
		{13},
	}
	result = spiralOrder(matrix)
	fmt.Println(result)
}
