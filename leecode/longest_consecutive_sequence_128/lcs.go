// Time        : 2019/07/24
// Description :

package longest_consecutive_sequence_128

// Given an unsorted array of integers,
// find the length of the longest consecutive elements sequence.
//
// Your algorithm should run in O(n) complexity.
//
// Example:
//
// Input: [100, 4, 200, 1, 3, 2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
// Therefore its length is 4.

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	nm := make(map[int]bool)
	for i := range nums {
		nm[nums[i]] = true
	}
	maxLength := 0
	for n := range nm {
		delete(nm, n)
		low, high := n, n
		length := 1
		for nm[low-1] || nm[high+1] {
			if nm[low-1] {
				delete(nm, low-1)
				low--
				length++
			}
			if nm[high+1] {
				delete(nm, high+1)
				high++
				length++
			}
		}
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
