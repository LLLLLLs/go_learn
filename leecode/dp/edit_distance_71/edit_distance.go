// Time        : 2019/07/09
// Description :

package edit_distance_71

// Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.
//
// You have the following 3 operations permitted on a word:
//
// Insert a character
// Delete a character
// Replace a character
// Example 1:
//
// Input: word1 = "horse", word2 = "ros"
// Output: 3
// Explanation:
// horse -> rorse (replace 'h' with 'r')
// rorse -> rose (remove 'r')
// rose -> ros (remove 'e')
// Example 2:
//
// Input: word1 = "intention", word2 = "execution"
// Output: 5
// Explanation:
// intention -> inention (remove 't')
// inention -> enention (replace 'i' with 'e')
// enention -> exention (replace 'n' with 'x')
// exention -> exection (replace 'n' with 'c')
// exection -> execution (insert 'u')

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word2)+1)
	for i := range dp {
		dp[i] = make([]int, len(word1)+1)
		dp[i][0] = i
	}
	for i := 0; i < len(word1)+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(word2)+1; i++ {
		for j := 1; j < len(word1)+1; j++ {
			if word2[i-1] == word1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				if dp[i][j-1] < dp[i-1][j] {
					dp[i][j] = dp[i][j-1] + 1
				}
			}
		}
	}
	return dp[len(word2)][len(word1)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
