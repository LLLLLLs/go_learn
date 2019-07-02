// Time        : 2019/01/24
// Description :

package wildcard_match_44

import (
	"fmt"
	"testing"
)

func TestWildcard(t *testing.T) {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("acdcb", "a*c?b"))
	fmt.Println(isMatch("adceb", "a*b"))
	fmt.Println(isMatch("a", "*a*"))
	fmt.Println(isMatch("b", "*a*"))
	fmt.Println(isMatch("bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab", "b*b*ab**ba*b**b***bba"))
}
