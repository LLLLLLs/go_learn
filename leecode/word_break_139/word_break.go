// Time        : 2019/07/26
// Description :

package word_break_139

// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
// determine if s can be segmented into a space-separated sequence of one or more dictionary words.
//
// Note:
//
// The same word in the dictionary may be reused multiple times in the segmentation.
// You may assume the dictionary does not contain duplicate words.
// Example 1:
//
// Input: s = "leetcode", wordDict = ["leet", "code"]
// Output: true
// Explanation: Return true because "leetcode" can be segmented as "leet code".
// Example 2:
//
// Input: s = "applepenapple", wordDict = ["apple", "pen"]
// Output: true
// Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
//              Note that you are allowed to reuse a dictionary word.
// Example 3:
//
// Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// Output: false

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]struct{})
	for i := range wordDict {
		words[wordDict[i]] = struct{}{}
	}
	match := make([]bool, len(s)+1)
	match[0] = true
	for i := range s {
		for k := i; k >= 0; k-- {
			if _, ok := words[s[k:i+1]]; ok && match[k] {
				match[i+1] = true
				break
			}
		}
	}
	return match[len(s)]
}
