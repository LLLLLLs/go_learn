// Time        : 2019/07/24
// Description :

package word_ladder_I_127

// Given two words (beginWord and endWord), and a dictionary's word list,
// find the length of shortest transformation sequence from beginWord to endWord, such that:
//
// Only one letter can be changed at a time.
// Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
// Note:
//
// Return 0 if there is no such transformation sequence.
// All words have the same length.
// All words contain only lowercase alphabetic characters.
// You may assume no duplicates in the word list.
// You may assume beginWord and endWord are non-empty and are not the same.
// Example 1:
//
// Input:
// beginWord = "hit",
// endWord = "cog",
// wordList = ["hot","dot","dog","lot","log","cog"]
//
// Output: 5
//
// Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
// return its length 5.
// Example 2:
//
// Input:
// beginWord = "hit"
// endWord = "cog"
// wordList = ["hot","dot","dog","lot","log"]
//
// Output: 0
//
// Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

func ladderLength(beginWord string, endWord string, wordList []string) int {
	words := make(map[string]struct{})
	for i := range wordList {
		words[wordList[i]] = struct{}{}
	}
	if _, ok := words[endWord]; !ok {
		return 0
	}
	begin, end := make(map[string]struct{}), make(map[string]struct{})
	begin[beginWord] = struct{}{}
	end[endWord] = struct{}{}
	length := 2
	back := false
	for len(begin) != 0 && len(end) != 0 {
		if len(begin) > len(end) {
			begin, end = end, begin
			back = !back
		}
		for w := range begin {
			delete(words, w)
		}
		newBegin := make(map[string]struct{})
		match := false
		for w := range begin {
			for i := range w {
				for j := 'a'; j <= 'z'; j++ {
					newWord := w[:i] + string(j) + w[i+1:]
					if _, ok := words[newWord]; ok {
						if _, ok2 := end[newWord]; ok2 {
							return length
						}
						newBegin[newWord] = struct{}{}
						match = true
					}
				}
			}
		}
		if match {
			length++
		}
		begin = newBegin
	}
	return 0
}
