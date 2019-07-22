// Time        : 2019/07/19
// Description :

package interleaving_string_97

//Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.
//
//Example 1:
//
//Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//Output: true
//Example 2:
//
//Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//Output: false

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	if s3 == "" {
		return true
	}
	dp := make([][]bool, len(s2)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s1)+1)
	}
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = s1[i-1] == s3[i-1]
		if !dp[0][i] {
			break
		}
	}
	for i := 1; i < len(dp); i++ {
		dp[i][0] = s2[i-1] == s3[i-1]
		if !dp[i][0] {
			break
		}
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = (dp[i-1][j] && (s2[i-1] == s3[i+j-1])) || (dp[i][j-1] && (s1[j-1] == s3[i+j-1]))
		}
	}
	return dp[len(s2)][len(s1)]
}
