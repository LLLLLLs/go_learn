// Time        : 2019/07/10
// Description :

package remove_duplicates_I_83

import (
	"golearn/leecode/linked_list/base"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	deleteDuplicates(base.NewList(1, 2, 2, 3, 3, 4, 5)).Print()
}
