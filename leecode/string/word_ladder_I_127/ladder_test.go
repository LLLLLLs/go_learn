// Time        : 2019/07/24
// Description :

package word_ladder_I_127

import (
	"fmt"
	"testing"
)

func TestLadder(t *testing.T) {
	fmt.Println(ladderLength(
		"hit",
		"cog",
		[]string{"hot", "dot", "dog", "lot", "log", "cog"}),
	)
}
