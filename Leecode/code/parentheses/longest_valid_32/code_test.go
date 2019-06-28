// Time        : 2019/06/28
// Description :

package longest_valid_32

import (
	"fmt"
	"testing"
)

func TestLongestValid(t *testing.T) {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("()(()"))
	fmt.Println(longestValidParentheses("(()(((()"))
}
