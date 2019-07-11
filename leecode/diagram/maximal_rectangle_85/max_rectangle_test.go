// Time        : 2019/07/11
// Description :

package maximal_rectangle_85

import (
	"fmt"
	"testing"
)

func TestMaxRectangle(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalRectangle(matrix))
	matrix = [][]byte{
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(maximalRectangle(matrix))
}
