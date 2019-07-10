// Time        : 2019/07/10
// Description :

package min_window_substring_76

// Given a string S and a string T,
// find the minimum window in S which will contain all the characters in T in complexity O(n).
//
// Example:
//
// Input: S = "ADOBECODEBANC", T = "ABC"
// Output: "BANC"
// Note:
//
// If there is no such window in S that covers all characters in T, return the empty string "".
// If there is such window, you are guaranteed that there will always be only one unique minimum window in S.

func minWindow(s string, t string) string {
	need := [128]int{}
	have := [128]int{}
	for i := range t {
		need[t[i]]++
	}
	var result string
	for i, j, count := 0, 0, 0; j < len(s); j++ {
		if have[s[j]] < need[s[j]] {
			count++
		}
		have[s[j]]++
		for i <= j && have[s[i]] > need[s[i]] {
			have[s[i]]--
			i++
		}
		if count == len(t) && (len(result) == 0 || j-i+1 < len(result)) {
			result = s[i : j+1]
		}
	}
	return result
}

func minWindow3(s string, t string) string {
	have := [128]int{}
	need := [128]int{}
	for i := range t {
		need[t[i]]++
	}

	size, total := len(s), len(t)

	min := size + 1
	res := ""

	for i, j, count := 0, 0, 0; j < size; j++ {
		if have[s[j]] < need[s[j]] {
			count++
		}
		have[s[j]]++

		for i <= j && have[s[i]] > need[s[i]] {
			have[s[i]]--
			i++
		}

		length := j - i + 1
		if count == total && min > length {
			min = length
			res = s[i : j+1]
		}
	}

	return res
}

func minWindow2(s string, t string) string {
	tMap := make(map[byte]int)
	for i := range t {
		tMap[t[i]]++
	}
	var result string
	left, right, curMap := 0, -1, make(map[byte]int)
	match := func() bool {
		for b, c := range tMap {
			if curMap[b] < c {
				return false
			}
		}
		return true
	}
	nodeList := make([]interface{}, 0, len(s))
	for i := range s {
		if tMap[s[i]] != 0 {
			nodeList = append(nodeList, i, s[i])
		}
	}
	for {
		if !match() {
			right++
			if right == len(nodeList)/2 {
				return result
			} else {
				curMap[nodeList[right*2+1].(byte)]++
			}
		} else {
			if len(result) == 0 || len(result) > nodeList[right*2].(int)-nodeList[left*2].(int)+1 {
				result = s[nodeList[left*2].(int) : nodeList[right*2].(int)+1]
			}
			if len(result) == len(t) {
				return result
			}
			curMap[nodeList[left*2+1].(byte)]--
			left++
		}
	}
}
