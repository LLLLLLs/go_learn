// Time        : 2019/07/03
// Description :

package solve_I_51

import (
	"golearn/util"
	"testing"
)

func TestNQueens(t *testing.T) {
	result := solveNQueens(4)
	util.Print2DimensionList(result)
	result = solveNQueens(8)
	util.Print2DimensionList(result)
}
