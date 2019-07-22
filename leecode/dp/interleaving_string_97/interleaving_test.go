// Time        : 2019/07/19
// Description :

package interleaving_string_97

import (
	"fmt"
	"testing"
)

func TestIsInterleaving(t *testing.T) {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbccc"))
	fmt.Println(isInterleave("db", "b", "cbb"))
}
