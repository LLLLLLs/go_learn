// Time        : 2019/07/10
// Description :

package remove_duplicates_II_82

import (
	"go_learn/leecode/linked_list/base"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	deleteDuplicates(base.NewList(1, 2, 3, 3, 4, 4, 5)).Print()
	deleteDuplicates(base.NewList(1, 2, 2)).Print()
}
