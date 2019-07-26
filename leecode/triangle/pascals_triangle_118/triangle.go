// Time        : 2019/07/23
// Description :

package pascals_triangle_118

// Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
//
//
// In Pascal's triangle, each number is the sum of the two numbers directly above it.
//
// Example:
//
// Input: 5
// Output:
// [
//      [1],
//     [1,1],
//    [1,2,1],
//   [1,3,3,1],
//  [1,4,6,4,1]
// ]

func generate(numRows int) [][]int {
	tri := make([][]int, numRows)
	for i := range tri {
		if i == 0 {
			tri[i] = []int{1}
			continue
		}
		mid := make([]int, i+1)
		for k := range mid {
			if k == 0 || k == len(mid)-1 {
				mid[k] = 1
				continue
			}
			mid[k] = tri[i-1][k-1] + tri[i-1][k]
		}
		tri[i] = mid
	}
	return tri
}
