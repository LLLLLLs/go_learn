// Time        : 2019/07/10
// Description :

package combination_77

// Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
//
// Example:
//
// Input: n = 4, k = 2
// Output:
// [
//   [2,4],
//   [3,4],
//   [2,3],
//   [1,2],
//   [1,3],
//   [1,4],
// ]

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	mid := make([]int, 0, k)
	recursive(1, n, k, mid, &result)
	return result
}

func recursive(begin, end, k int, mid []int, result *[][]int) {
	if end-begin+1 < k {
		return
	}
	if k == 0 {
		*result = append(*result, mid)
		return
	}
	for i := begin; i <= end; i++ {
		tmp := make([]int, len(mid), cap(mid))
		copy(tmp, mid)
		tmp = append(tmp, i)
		recursive(i+1, end, k-1, tmp, result)
	}
}

func combineBinary(n int, k int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < (1 << uint(n)); i++ {
		tmp := make([]int, 0, k)
		for j := 0; j < n; j++ {
			if i&(1<<uint(j)) != 0 {
				tmp = append(tmp, j+1)
			}
		}
		if len(tmp) == k {
			result = append(result, tmp)
		}
	}
	return result
}
