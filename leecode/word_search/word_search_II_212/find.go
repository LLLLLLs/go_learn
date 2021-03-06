// Time        : 2019/07/10
// Description :

package word_search_II_212

// Given a 2D board and a list of words from the dictionary, find all words in the board.
//
// Each word must be constructed from letters of sequentially adjacent cell,
// where "adjacent" cells are those horizontally or vertically neighboring.
// The same letter cell may not be used more than once in a word.
//
//
//
// Example:
//
// Input:
// board = [
//   ['o','a','a','n'],
//   ['e','t','a','e'],
//   ['i','h','k','r'],
//   ['i','f','l','v']
// ]
// words = ["oath","pea","eat","rain"]
//
// Output: ["eat","oath"]

//func findWords(board [][]byte, words []string) []string {
//	path := make([][]bool, len(board))
//	for i := range path {
//		path[i] = make([]bool, len(board[i]))
//	}
//	var remain = make([]string, len(words))
//	copy(remain, words)
//	for i := 0; i < len(board); i++ {
//		for j := 0; j < len(board[i]); j++ {
//			wordM := make(map[string]struct{})
//			for k := 0; k < len(remain); k++ {
//				if remain[k][0] == board[i][j] {
//					wordM[remain[k]] = struct{}{}
//				}
//			}
//			path[i][j] = true
//			backtrack(board, path, i, j, 1, wordM, &remain)
//			path[i][j] = false
//		}
//	}
//	var result = make([]string, len(words)-len(remain))
//	var j, k = 0, 0
//	for i := range words {
//		if j == len(remain) || words[i] != remain[j] {
//			result[k] = words[i]
//			k++
//		} else {
//			j++
//		}
//	}
//	return result
//}
//
//var di = []int{0, 1, 0, -1}
//var dj = []int{1, 0, -1, 0}
//
//func backtrack(board [][]byte, path [][]bool, i, j, index int, words map[string]struct{}, remain *[]string) {
//	for k := 0; k < len(*remain); {
//		word := (*remain)[k]
//		if _, ok := words[word]; ok && index == len(word) {
//			delete(words, word)
//			*remain = append((*remain)[:k], (*remain)[k+1:]...)
//			continue
//		}
//		k++
//	}
//	if len(words) == 0 {
//		return
//	}
//	for k := 0; k < 4; k++ {
//		ni, nj := i+di[k], j+dj[k]
//		if ni >= 0 && nj >= 0 && ni < len(board) && nj < len(board[ni]) && !path[ni][nj] {
//			tmp := make(map[string]struct{})
//			for i := range *remain {
//				word := (*remain)[i]
//				if _, ok := words[word]; ok && board[ni][nj] == word[index] {
//					tmp[word] = struct{}{}
//				}
//			}
//			path[ni][nj] = true
//			backtrack(board, path, ni, nj, index+1, tmp, remain)
//			path[ni][nj] = false
//			if len(*remain) == 0 {
//				return
//			}
//		}
//	}
//	return
//}

func findWords(board [][]byte, words []string) []string {
	path := make([][]bool, len(board))
	for i := range path {
		path[i] = make([]bool, len(board[i]))
	}
	var result = make([]string, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			for k := 0; k < len(words); {
				if board[i][j] == words[k][0] {
					path[i][j] = true
					if backtrack(board, path, i, j, words[k][1:]) {
						result = append(result, words[k])
						words = append(words[:k], words[k+1:]...)
					} else {
						k++
					}
					path[i][j] = false
				} else {
					k++
				}
			}
		}
	}
	return result
}

var di = []int{0, 1, 0, -1}
var dj = []int{1, 0, -1, 0}

func backtrack(board [][]byte, path [][]bool, i, j int, word string) bool {
	if len(word) == 0 {
		return true
	}
	for k := 0; k < 4; k++ {
		ni, nj := i+di[k], j+dj[k]
		if ni >= 0 && nj >= 0 && ni < len(board) && nj < len(board[ni]) && !path[ni][nj] && board[ni][nj] == word[0] {
			path[ni][nj] = true
			if backtrack(board, path, ni, nj, word[1:]) {
				path[ni][nj] = false
				return true
			}
			path[ni][nj] = false
		}
	}
	return false
}
