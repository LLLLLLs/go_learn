// Time        : 2019/07/09
// Description :

package textjustification_68

// Given an array of words and a width maxWidth,
// format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.
//
// You should pack your words in a greedy approach; that is, pack as many words as you can in each line.
// Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.
//
// Extra spaces between words should be distributed as evenly as possible.
// If the number of spaces on a line do not divide evenly between words,
// the empty slots on the left will be assigned more spaces than the slots on the right.
//
// For the last line of text, it should be left justified and no extra space is inserted between words.
//
// Note:
//
// A word is defined as a character sequence consisting of non-space characters only.
// Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
// The input array words contains at least one word.
// Example 1:
//
// Input:
// words = ["This", "is", "an", "example", "of", "text", "justification."]
// maxWidth = 16
// Output:
// [
//    "This    is    an",
//    "example  of text",
//    "justification.  "
// ]
// Example 2:
//
// Input:
// words = ["What","must","be","acknowledgment","shall","be"]
// maxWidth = 16
// Output:
// [
//   "What   must   be",
//   "acknowledgment  ",
//   "shall be        "
// ]
// Explanation: Note that the last line is "shall be    " instead of "shall     be",
//              because the last line must be left-justified instead of fully-justified.
//              Note that the second line is also left-justified becase it contains only one word.
// Example 3:
//
// Input:
// words = ["Science","is","what","we","understand","well","enough","to","explain",
//          "to","a","computer.","Art","is","everything","else","we","do"]
// maxWidth = 20
// Output:
// [
//   "Science  is  what we",
//   "understand      well",
//   "enough to explain to",
//   "a  computer.  Art is",
//   "everything  else  we",
//   "do                  "
// ]

func fullJustify(words []string, maxWidth int) []string {
	var result = make([]string, 0)
	for i := 0; i < len(words); {
		width := 0
		j := i
		for ; ; j++ {
			width += len(words[j]) + 1
			if j+1 > len(words)-1 || width+len(words[j+1]) > maxWidth {
				width--
				break
			}
		}
		var str string
		count := j - i
		remain := maxWidth - width
		if j == len(words)-1 {
			for k := i; k <= j; k++ {
				str += words[k] + " "
			}
			str = str[:len(str)-1] + genSpace(remain)
			result = append(result, str)
			break
		}
		for k := 0; k < count; k++ {
			var space = remain / count
			if k < remain%count {
				space++
			}
			str += words[i+k] + genSpace(space+1)
		}
		str += words[j]
		if count == 0 {
			str += genSpace(remain)
		}
		result = append(result, str)
		i = j + 1
	}
	return result
}

func genSpace(n int) string {
	var str string
	for i := 0; i < n; i++ {
		str += " "
	}
	return str
}
