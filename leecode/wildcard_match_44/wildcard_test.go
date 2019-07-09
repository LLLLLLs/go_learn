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
	//fmt.Println(isMatch("bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab", "b*b*ab**ba*b**b***bba"))
}

func TestWildMatchDp(t *testing.T) {
	fmt.Println(isMatchDp("aa", "a"))
	fmt.Println(isMatchDp("acdcb", "a*c?b"))
	fmt.Println(isMatchDp("adceb", "a*b"))
	fmt.Println(isMatchDp("a", "*a*"))
	fmt.Println(isMatchDp("b", "*a*"))
	fmt.Println(isMatchDp("bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab", "b*b*ab**ba*b**b***bba"))
}
