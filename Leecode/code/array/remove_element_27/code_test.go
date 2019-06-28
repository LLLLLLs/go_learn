// Time        : 2019/06/26
// Description :

package remove_element_27

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	list := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(list, val))
	fmt.Print(list)
}
