// Time        : 2019/06/26
// Description :

package swap_nodes_24

import (
	"go_learn/leecode/linked_list/base"
	"testing"
)

func TestSwap(t *testing.T) {
	l := base.NewList(1, 2, 3, 4, 5)
	swapPairs(l).Print()
}
