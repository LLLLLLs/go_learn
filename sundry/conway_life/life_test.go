// Time        : 2019/11/11
// Description :

package conwaylife

import (
	"fmt"
	"testing"
	"time"
)

func TestNextStep(t *testing.T) {
	board := [][]bool{
		{false, true, true, false},
		{true, true, false, true},
		{true, false, true, false},
		{false, true, false, false},
	}
	//board := [][]bool{
	//	{true, false},
	//	{true, true},
	//	{true, true},
	//}
	board = formatBoard(board)
	for len(board) != 0 {
		printBoard(board)
		board = NextStep(board)
		time.Sleep(time.Millisecond * 300)
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
