// Time        : 2019/07/26
// Description :

package word_break_139

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
