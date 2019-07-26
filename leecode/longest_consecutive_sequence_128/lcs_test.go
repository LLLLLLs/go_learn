// Time        : 2019/07/24
// Description :

package longest_consecutive_sequence_128

import (
	"fmt"
	"testing"
)

func TestLcs(t *testing.T) {
	fmt.Println(longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
