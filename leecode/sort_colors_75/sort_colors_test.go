// Time        : 2019/07/10
// Description :

package sort_colors_75

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	colors := []int{2, 0, 2, 1, 1, 0}
	sortColors(colors)
	fmt.Println(colors)
	colors = []int{0, 2, 0}
	sortColors(colors)
	fmt.Println(colors)
}
