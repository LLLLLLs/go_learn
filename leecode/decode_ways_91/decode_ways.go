// Time        : 2019/07/12
// Description :

package decode_ways_91

// A message containing letters from A-Z is being encoded to numbers using the following mapping:
//
// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// Given a non-empty string containing only digits, determine the total number of ways to decode it.
//
// Example 1:
//
// Input: "12"
// Output: 2
// Explanation: It could be decoded as "AB" (1 2) or "L" (12).
// Example 2:
//
// Input: "226"
// Output: 3
// Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

func numDecodings(s string) int {
	if s == "" || s[0] == '0' {
		return 0
	}
	f := [2]int{1, 1}
	if len(s) < 2 {
		return f[len(s)]
	}
	for i := 1; i < len(s); i++ {
		double := s[i-1 : i+1]
		if s[i] == '0' && double > "26" || double == "00" {
			return 0
		} else if s[i] == '0' {
			f[1] = f[0]
			continue
		}
		if double <= "26" && s[i-1] != '0' {
			f[1], f[0] = f[0]+f[1], f[1]
		} else {
			f[0] = f[1]
		}
	}
	return f[1]
}
