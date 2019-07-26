// Time        : 2019/07/25
// Description :

package palindrome_partitioning_131

//Given a string s, partition s such that every substring of the partition is a palindrome.
//
//Return all possible palindrome partitioning of s.
//
//Example:
//
//Input: "aab"
//Output:
//[
//  ["aa","b"],
//  ["a","a","b"]
//]

func partition(s string) [][]string {
	length := len(s)
	if length < 2 {
		return [][]string{{s}}
	}
	dp := make([][]bool, length+1)
	for i := range dp {
		dp[i] = make([]bool, i+1)
	}
	for i := 0; i < length; i++ {
		dp[i][i] = true
		dp[i+1][i] = true
	}
	dp[length][length] = true
	for i := 2; i < len(dp); i++ {
		for j := 0; j < len(dp[i])-2; j++ {
			dp[i][j] = s[j] == s[i-1] && dp[i-1][j+1]
		}
	}
	result := make([][]string, 0)
	stack := make([]string, 0)
	var generate func(start int)
	generate = func(start int) {
		if start == length {
			result = append(result, append([]string{}, stack...))
			return
		}
		for i := start + 1; i < len(dp); i++ {
			if dp[i][start] {
				stack = append(stack, s[start:i])
				generate(i)
				stack = stack[:len(stack)-1]
			}
		}
	}
	generate(0)
	return result
}
