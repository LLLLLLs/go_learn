// Time        : 2019/07/03
// Description :

package permutations_II_47

import (
	"fmt"
	"golearn/utils"
	"testing"
)

func TestPermutation(t *testing.T) {
	result := permuteUnique([]int{1, 1, 3})
	utils.Print2DimensionList(result)
	result = permuteUnique([]int{-1, 2, -1, 2, 1, -1, 2, 1})
	utils.Print2DimensionList(result)
	fmt.Println(len(result))
}
