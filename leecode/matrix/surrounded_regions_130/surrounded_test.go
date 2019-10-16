// Time        : 2019/07/25
// Description :

package surrounded_regions_130

import (
	"golearn/util"
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
	util.Print2DimensionList(board)
	board = [][]byte{
		{'O'},
	}
	solve(board)
	util.Print2DimensionList(board)
}
