// Time        : 2019/07/11
// Description :

package scramble_string_87

import (
	"fmt"
	"testing"
)

func TestScramble(t *testing.T) {
	fmt.Println(isScramble("great", "rgeat"))
	fmt.Println(isScramble("abcde", "caebd"))
	fmt.Println(isScramble("abc", "bac"))
}
