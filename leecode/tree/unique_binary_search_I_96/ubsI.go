// Time        : 2019/07/19
// Description :

package unique_binary_search_I_96

//Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?
//
//Example:
//
//Input: 3
//Output: 5
//Explanation:
//Given n = 3, there are a total of 5 unique BST's:
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3

func numTrees(n int) int {
	if n <= 1 {
		return n
	}
	var f = make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i < len(f); i++ {
		for k := 0; k < i; k++ {
			f[i] += f[k] * f[i-k-1]
		}
	}
	return f[n]
}
