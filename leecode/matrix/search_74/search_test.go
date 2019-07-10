// Time        : 2019/07/10
// Description :

package search_74

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(matrix, 3))

	matrix = [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(matrix, 13))

	matrix = [][]int{
		{1},
	}
	fmt.Println(searchMatrix(matrix, 2))
}
