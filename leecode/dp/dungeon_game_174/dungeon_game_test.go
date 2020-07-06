//@author: lls
//@time: 2020/07/03
//@desc:

package dungeon_game_174

import (
	"fmt"
	"testing"
)

func TestDungeonGame(t *testing.T) {
	// -2 (K)	-3		3
	// -5		-10		1
	// 10		30		-5 (P)
	fmt.Println(calculateMinimumHP([][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}))
	// [3,-20,30],
	// [-3,4,0]
	fmt.Println(calculateMinimumHP([][]int{
		{3, -20, 30},
		{-3, 4, 0},
	}))

	// [1,-3,3],
	// [0,-2,0],
	// [-3,-3,-3]
	fmt.Println(calculateMinimumHP([][]int{
		{1, -3, 3},
		{0, -2, 0},
		{-3, -3, -3},
	}))
}
