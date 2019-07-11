// Time        : 2019/07/11
// Description :

package partition_86

import (
	"go_learn/leecode/linked_list/base"
	"testing"
)

func TestPartition(t *testing.T) {
	partition(base.NewList(1, 4, 3, 2, 5, 2), 3).Print()
	partition(base.NewList(3, 1, 2), 3).Print()
}
