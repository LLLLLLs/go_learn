// Time        : 2019/07/04
// Description :

package spiral_II_59

import (
	"go_learn/utils"
	"testing"
)

func TestSpiral(t *testing.T) {
	result := generateMatrix(3)
	utils.Print2DimensionList(result)
}
