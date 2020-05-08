// Time        : 2019/11/11
// Description :

package conwaylife

import (
	"fmt"
	"testing"
	"time"
)

func TestNextStep(t *testing.T) {
	// . X X .
	// X X . X
	// X . X .
	// . X . .
	//board := [][]bool{
	//	{false, true, true, false},
	//	{true, true, false, true},
	//	{true, false, true, false},
	//	{false, true, false, false},
	//}
	// X .
	// X X
	// X X
	board := [][]bool{
		{true, false},
		{true, true},
		{true, true},
	}
	board = formatBoard(board)
	printBoard(board)
	for len(board) != 0 {
		board = NextStep(board)
		time.Sleep(time.Millisecond * 300)
		printBoard(board)
	}
}

func printBoard(board [][]bool) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] {
				fmt.Printf("X ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func BenchmarkNextStep(b *testing.B) {
	board := [][]bool{
		{false, true, true, false},
		{true, true, false, true},
		{true, false, true, false},
		{false, true, false, false},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpy := formatBoard(board)
		for len(cpy) != 0 {
			cpy = NextStep(cpy)
		}
	}
}
