// Time        : 2019/07/03
// Description :

package permutations_I_46

import (
	"go_learn/utils"
	"testing"
)

func TestPermutation(t *testing.T) {
	result := permute([]int{1, 2, 3})
	utils.Print2DimensionList(result)
}
