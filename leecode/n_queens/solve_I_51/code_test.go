// Time        : 2019/07/03
// Description :

package solve_I_51

import (
	"go_learn/utils"
	"testing"
)

func TestNQueens(t *testing.T) {
	result := solveNQueens(4)
	utils.Print2DimensionList(result)
	result = solveNQueens(8)
	utils.Print2DimensionList(result)
}
