// Time        : 2019/06/26
// Description :

package remove_duplicates_26

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	list := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(list))
}
