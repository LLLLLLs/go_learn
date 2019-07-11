// Time        : 2019/07/10
// Description :

package word_search_II_212

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words))

	board = [][]byte{
		{'a', 'b'},
		{'a', 'a'},
	}
	words = []string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"}
	fmt.Println(findWords(board, words))

	board = [][]byte{
		{'a', 'b'},
	}
	words = []string{"a", "b"}
	fmt.Println(findWords(board, words))
}
