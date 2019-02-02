// Time        : 2019/01/24
// Description :

package wildcard_match

import (
	"fmt"
	"testing"
)

func TestWildcard(t *testing.T) {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("acdcb", "a*c?b"))
	fmt.Println(isMatch("adceb", "a*b"))
}
