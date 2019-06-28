// Time        : 2019/06/26
// Description :

package generate_22

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
// For example, given n = 3, a solution set is:
//
// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]

func generateParenthesis1(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	list1 := generateParenthesis1(n - 1)
	var result = make(map[string]bool)
	for i := range list1 {
		for _, s := range addP([]rune(list1[i])) {
			result[s] = true
		}
	}
	var list2 = make([]string, 0, len(result))
	for s := range result {
		list2 = append(list2, s)
	}
	return list2
}

func addP(rs []rune) []string {
	rs = append([]rune{'(', ')'}, rs...)
	result := []string{string(rs)}
	for i := 1; i < len(rs)-1; i++ {
		rs[i], rs[i+1] = rs[i+1], rs[i]
		result = append(result, string(rs))
	}
	return result
}

func generateParenthesis2(n int) []string {
	result := make([]string, 0)
	mid := make([]byte, 0)
	backtrack(&result, mid, 0, 0, n)
	return result
}

func backtrack(result *[]string, mid []byte, open, close, max int) {
	if len(mid) == max*2 {
		*result = append(*result, string(mid))
		return
	}
	if open < max {
		backtrack(result, append(mid, '('), open+1, close, max)
	}
	if close < open {
		backtrack(result, append(mid, ')'), open, close+1, max)
	}
}
