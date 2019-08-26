// Time        : 2019/07/31
// Description :

package reverse_151

import (
	"strings"
)

// Given an input string, reverse the string word by word.
//
// Example 1:
//
// Input: "the sky is blue"
// Output: "blue is sky the"
// Example 2:
//
// Input: "  hello world!  "
// Output: "world! hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.
// Example 3:
//
// Input: "a good   example"
// Output: "example good a"
// Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
//
// Note:
//
// A word is defined as a sequence of non-space characters.
// Input string may contain leading or trailing spaces. However, your reversed string should not contain leading or trailing spaces.
// You need to reduce multiple spaces between two words to a single space in the reversed string.

func reverseWords(s string) string {
	ss := strings.Fields(strings.TrimSpace(s))
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[len(ss)-1-i] = ss[len(ss)-1-i], ss[i]
	}
	return strings.Join(ss, " ")
}
