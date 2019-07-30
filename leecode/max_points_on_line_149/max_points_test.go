// Time        : 2019/07/30
// Description :

package max_points_on_line_149

import (
	"fmt"
	"testing"
)

func TestMaxPoints(t *testing.T) {
	points := [][]int{
		{1, 1},
		{3, 2},
		{5, 3},
		{4, 1},
		{2, 3},
		{1, 4},
	}
	fmt.Println(maxPoints(points))
	points = [][]int{
		{1, 1},
		{1, -1},
		{0, 0},
	}
	fmt.Println(maxPoints(points))
}
