// Time        : 2019/07/04
// Description :

package unique_paths_II_63

import (
	"fmt"
	"testing"
)

func TestUniquePath(t *testing.T) {
	grid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(grid))
}
