// Time        : 2019/07/26
// Description :

package word_break_II_140

// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
// add spaces in s to construct a sentence where each word is a valid dictionary word.
// Return all such possible sentences.
//
// Note:
//
// The same word in the dictionary may be reused multiple times in the segmentation.
// You may assume the dictionary does not contain duplicate words.
// Example 1:
//
// Input:
// s = "catsanddog"
// wordDict = ["cat", "cats", "and", "sand", "dog"]
// Output:
// [
//   "cats and dog",
//   "cat sand dog"
// ]
// Example 2:
//
// Input:
// s = "pineapplepenapple"
// wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
// Output:
// [
//   "pine apple pen apple",
//   "pineapple pen apple",
//   "pine applepen apple"
// ]
// Explanation: Note that you are allowed to reuse a dictionary word.
// Example 3:
//
// Input:
// s = "catsandog"
// wordDict = ["cats", "dog", "sand", "and", "cat"]
// Output:
// []

func wordBreak(s string, wordDict []string) []string {
	return match(s, wordDict, map[string][]string{})
}

// 缓存中间值
func match(s string, words []string, cache map[string][]string) []string {
	if res, ok := cache[s]; ok {
		return res
	}
	if s == "" {
		return []string{""}
	}
	res := make([]string, 0)
	for i := range words {
		if len(words[i]) <= len(s) && words[i] == s[:len(words[i])] {
			for _, w := range match(s[len(words[i]):], words, cache) {
				if w == "" {
					res = append(res, words[i])
				} else {
					res = append(res, words[i]+" "+w)
				}
			}
		}
	}
	cache[s] = res
	return res
}
