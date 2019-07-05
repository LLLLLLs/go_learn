// Time        : 2019/01/24
// Description :

package jump_game_II_45

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	fmt.Println(jump([]int{2, 3, 1, 1, 4, 3, 1, 1, 5, 1, 1, 2, 3, 4, 1, 1}))
}

func TestJump2(t *testing.T) {
	fmt.Println(jump2([]int{2, 3, 1, 1, 4, 3, 1, 1, 5, 1, 1, 2, 3, 4, 1, 1}))
}

func BenchmarkJump1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jump([]int{2, 3, 1, 1, 4, 3, 1, 1, 5, 1, 1, 2, 3, 4, 1, 1})
	}
}

func BenchmarkJump2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jump2([]int{2, 3, 1, 1, 4, 3, 1, 1, 5, 1, 1, 2, 3, 4, 1, 1})
	}
}
