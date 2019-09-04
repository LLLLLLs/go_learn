// Time        : 2019/07/12
// Description :

package reverse_II_92

import (
	"golearn/leecode/linked_list/base"
	"testing"
)

func TestReverse(t *testing.T) {
	list := base.NewList(1, 2, 3, 4, 5)
	reverseBetween(list, 2, 4).Print()
}
