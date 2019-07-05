// Time        : 2019/07/04
// Description :

package length_of_last_word_58

// Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.
//
// If the last word does not exist, return 0.
//
// Note: A word is defined as a character sequence consists of non-space characters only.
//
// Example:
//
// Input: "Hello World"
// Output: 5

func lengthOfLastWord(s string) int {
	var result, empty = 0, true
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			empty = false
			result++
		} else {
			if !empty {
				break
			}
		}
	}
	return result
}
