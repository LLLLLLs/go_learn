// Time        : 2019/07/23
// Description :

package minimum_from_top_to_botton_120

// Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.
//
// For example, given the following triangle
//
// [
//      [2],
//     [3,4],
//    [6,5,7],
//   [4,1,8,3]
// ]
// The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
//
// Note:
//
// Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	for i := 1; i < len(triangle); i++ {
		for k := range triangle[i] {
			if k == 0 {
				triangle[i][k] += triangle[i-1][0]
				continue
			}
			if k == len(triangle[i])-1 {
				triangle[i][k] += triangle[i-1][len(triangle[i-1])-1]
				continue
			}
			triangle[i][k] += min(triangle[i-1][k-1], triangle[i-1][k])
		}
	}
	minimum := triangle[len(triangle)-1][0]
	for i := range triangle[len(triangle)-1] {
		if triangle[len(triangle)-1][i] < minimum {
			minimum = triangle[len(triangle)-1][i]
		}
	}
	return minimum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
