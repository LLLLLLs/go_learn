// Time        : 2019/07/23
// Description :

package pascals_triangle_II_119

//Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.
//
//Note that the row index starts from 0.
//
//In Pascal's triangle, each number is the sum of the two numbers directly above it.
//
//Example:
//
//Input: 3
//Output: [1,3,3,1]

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	last := []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		cur := make([]int, i+1)
		for k := range cur {
			if k == 0 || k == len(cur)-1 {
				cur[k] = 1
				continue
			}
			cur[k] = last[k-1] + last[k]
		}
		last = cur
	}
	return last
}
