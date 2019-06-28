// Time        : 2019/06/28
// Description :

package substr_concatenation_all_words_30

// You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.
//
// Example 1:
//
// Input:
//   s = "barfoothefoobarman",
//   words = ["foo","bar"]
// Output: [0,9]
// Explanation: Substrings starting at index 0 and 9 are "barfoor" and "foobar" respectively.
// The output order does not matter, returning [9,0] is fine too.
// Example 2:
//
// Input:
//   s = "wordgoodgoodgoodbestword",
//   words = ["word","good","best","word"]
// Output: []

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	length := len(words[0])
	totalLen := length * len(words)
	var init = make(map[string]int)
	var loop = make(map[string]int)
	for i := range words {
		init[words[i]]++
		loop[words[i]]++
	}
	reset := func() {
		for k := range init {
			loop[k] = init[k]
		}
	}
	var result = make([]int, 0)
	for i := 0; i <= len(s)-totalLen; i++ {
		var j = i
		var modify bool
		for {
			count, has := loop[s[j:j+length]]
			if !has || count == 0 {
				if modify {
					reset()
				}
				break
			} else {
				loop[s[j:j+length]]--
				modify = true
			}
			j += length
			if j-i == totalLen {
				result = append(result, i)
				reset()
				break
			}
		}
	}
	return result
}
