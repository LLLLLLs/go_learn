// Time        : 2019/07/25
// Description :

package surrounded_regions_130

import (
	"golearn/utils"
	"testing"
)

func TestSurrounded(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	utils.Print2DimensionList(board)
	board = [][]byte{
		{'O'},
	}
	solve(board)
	utils.Print2DimensionList(board)
}
