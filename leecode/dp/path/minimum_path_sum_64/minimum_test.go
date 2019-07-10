// Time        : 2019/07/04
// Description :

package minimum_path_sum_64

import (
	"fmt"
	"testing"
)

func TestMinimum(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}
