// Time        : 2019/01/24
// Description :

package wildcard_match_44

// Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*'.
//
// '?' Matches any single character.
// '*' Matches any sequence of characters (including the empty sequence).
// The matching should cover the entire input string (not partial).
//
// Note:
//
// s could be empty and contains only lowercase letters a-z.
// p could be empty and contains only lowercase letters a-z, and characters like ? or *.
//
// Example 1:
//
// Input:
// s = "aa"
// p = "a"
// Output: false
// Explanation: "a" does not match the entire string "aa".
// Example 2:
//
// Input:
// s = "aa"
// p = "*"
// Output: true
// Explanation: '*' matches any sequence.
// Example 3:
//
// Input:
// s = "cb"
// p = "?a"
// Output: false
// Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.
// Example 4:
//
// Input:
// s = "adceb"
// p = "*a*b"
// Output: true
// Explanation: The first '*' matches the empty sequence, while the second '*' matches the substring "dce".
// Example 5:
//
// Input:
// s = "acdcb"
// p = "a*c?b"
// Output: false

func isMatch(s string, p string) bool {
	if p == "*" {
		return true
	}
	if remain(p) > len(s) {
		return false
	}
	var i = 0
	for ; i < len(p); i++ {
		if p[i] == '?' {
			continue
		}
		if p[i] == '*' {
			if len(s)-i < remain(p[i+1:]) {
				return false
			}
			for j := len(s) - i; j >= 0; j-- {
				if isMatch(s[i+j:], p[i+1:]) {
					return true
				}
			}
			return false
		}
		if i >= len(s) || p[i] != s[i] {
			return false
		}
	}
	if i == len(s) {
		return true
	}
	return false
}

func remain(p string) int {
	var remain = 0
	for i := range p {
		if p[i] != '*' {
			remain++
		}
	}
	return remain
}

func isMatchDp(s string, p string) bool {
	var dp = make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true
	for j := len(p) - 1; j >= 0; j-- {
		if p[j] == '*' {
			dp[len(s)][j] = dp[len(s)][j+1]
		} else {
			dp[len(s)][j] = false
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			if p[j] == '*' {
				dp[i][j] = dp[i+1][j] || dp[i][j+1]
			} else if p[j] == '?' || p[j] == s[i] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[0][0]
}
