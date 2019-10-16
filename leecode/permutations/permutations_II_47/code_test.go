// Time        : 2019/07/03
// Description :

package permutations_II_47

import (
	"fmt"
	"golearn/util"
	"testing"
)

func TestPermutation(t *testing.T) {
	result := permuteUnique([]int{1, 1, 3})
	util.Print2DimensionList(result)
	result = permuteUnique([]int{-1, 2, -1, 2, 1, -1, 2, 1})
	util.Print2DimensionList(result)
	fmt.Println(len(result))
}
