// Time        : 2019/07/09
// Description :

package textjustification_68

import (
	"fmt"
	"testing"
)

func TestFullJustify(t *testing.T) {
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
		"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	result := fullJustify(words, 20)
	printResult(result)

	words = []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	result = fullJustify(words, 16)
	printResult(result)

	words = []string{"a"}
	result = fullJustify(words, 1)
	printResult(result)

	words = []string{"My", "momma", "always", "said,", "\"Life", "was", "like", "a", "box", "of", "chocolates.", "You", "never", "know", "what", "you're", "gonna", "get."}
	result = fullJustify(words, 20)
	printResult(result)
}

func printResult(r []string) {
	for i := range r {
		fmt.Println(r[i])
	}
	fmt.Println()
}
