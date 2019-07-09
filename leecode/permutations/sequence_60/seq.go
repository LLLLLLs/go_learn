// Time        : 2019/07/04
// Description :

package sequence_60

// The set [1,2,3,...,n] contains a total of n! unique permutations.
//
// By listing and labeling all of the permutations in order, we get the following sequence for n = 3:
//
// "123"
// "132"
// "213"
// "231"
// "312"
// "321"
// Given n and k, return the kth permutation sequence.
//
// Note:
//
// Given n will be between 1 and 9 inclusive.
// Given k will be between 1 and n! inclusive.
// Example 1:
//
// Input: n = 3, k = 3
// Output: "213"
// Example 2:
//
// Input: n = 4, k = 9
// Output: "2314"

func getPermutation(n int, k int) string {
	list := make([]byte, n)
	for i := byte(1); i <= byte(n); i++ {
		list[i-1] = i + '0'
	}
	result := recursive(list, k)
	return string(result)
}

var nStepMap = map[int]int{1: 1}

func nStep(n int) int {
	if nStepMap[n] != 0 {
		return nStepMap[n]
	}
	return n * nStep(n-1)
}

func recursive(list []byte, k int) []byte {
	if k == 0 || k == 1 {
		return list
	}
	var step = 1
	for nStep(step) < k {
		step++
	}
	reserve := append([]byte{}, list[:len(list)-step]...)
	next := append([]byte{}, list[len(list)-step:]...)
	choose := (k - 1) / nStep(step-1)
	if choose != 0 {
		reserve = append(reserve, next[choose])
		next = append(next[:choose], next[choose+1:]...)
	}
	return append(reserve, recursive(next, k-choose*nStep(step-1))...)
}
