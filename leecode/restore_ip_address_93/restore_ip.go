// Time        : 2019/07/12
// Description :

package restore_ip_address_93

import "strings"

// Given a string containing only digits, restore it by returning all possible valid IP address combinations.
//
// Example:
//
// Input: "25525511135"
// Output: ["255.255.11.135", "255.255.111.35"]

// 节省空间的小技巧：创建mid变量时指定其cap=4，
// 这样在递归回溯的过程中mid不会因为append而开辟新空间，
// 所以每个递归中mid指向的都是同一个地址，提高内存空间的利用率
// benchmark对比：
// 使用cap=0的mid，在每个递归中都创建一个新的数组存放mid ==> tmp := append(append([]string,mid...),str):
// 		BenchmarkRestoreIp-12	500000		2445 ns/op		1568 B/op		37 allocs/op
// 使用cap=4的mid，每个递归中使用同一个mid ==> mid = append(mid, s[:i+1])
//											backtrack(s[i+1:], mid, result)
//											mid = mid[:len(mid)-1]
// 		BenchmarkRestoreIp-12	5000000		342 ns/op		80 B/op			4 allocs/op
// 性能差了10倍~
func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	mid := make([]string, 0, 4)
	backtrack(s, mid, &result)
	return result
}

func backtrack(s string, mid []string, result *[]string) {
	if len(s) > 3*(4-len(mid)) {
		return
	}
	if len(s) == 0 && len(mid) == 4 {
		*result = append(*result, strings.Join(mid, "."))
		return
	}
	for i := 0; i < 3 && i < len(s); i++ {
		if (i < 2 || s[:i+1] <= "255") && !(i != 0 && s[0] == '0') {
			mid = append(mid, s[:i+1])
			backtrack(s[i+1:], mid, result)
			mid = mid[:len(mid)-1]
		}
	}
}
