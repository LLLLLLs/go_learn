// Time        : 2019/07/26
// Description :

package word_break_II_140

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	result := wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"})
	for i := range result {
		fmt.Println(result[i])
	}
}
