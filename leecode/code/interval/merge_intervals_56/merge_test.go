// Time        : 2019/07/04
// Description :

package merge_intervals_56

import (
	"fmt"
	"testing"
)

func TestMergeInterval(t *testing.T) {
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	output := merge(input)
	fmt.Println(output)
}
