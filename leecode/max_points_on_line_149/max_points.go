// Time        : 2019/07/30
// Description :

package max_points_on_line_149

// Given n points on a 2D plane, find the maximum number of points that lie on the same straight line.
//
// Example 1:
//
// Input: [[1,1],[2,2],[3,3]]
// Output: 3
// Explanation:
// ^
// |
// |        o
// |     o
// |  o
// +------------->
// 0  1  2  3  4
// Example 2:
//
// Input: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
// Output: 4
// Explanation:
// ^
// |
// |  o
// |     o        o
// |        o
// |  o        o
// +------------------->
// 0  1  2  3  4  5  6

func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}
	var max int
	var verticalMap = make(map[int]struct{})
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			var vertical bool
			var count = 2
			if points[i][0] == points[j][0] {
				vertical = true
				if _, ok := verticalMap[points[i][0]]; ok {
					continue
				} else {
					verticalMap[points[i][0]] = struct{}{}
				}
			}
			for k := 0; k < len(points); k++ {
				if k == i || k == j {
					continue
				}
				if vertical {
					if points[k][0] == points[i][0] {
						count++
					}
					continue
				}
				if (points[i][0]-points[j][0])*(points[i][1]-points[k][1]) ==
					(points[i][0]-points[k][0])*(points[i][1]-points[j][1]) {
					if _, ok := verticalMap[points[k][0]]; !ok && k < j {
						break
					}
					count++
				}
			}
			if count > max {
				max = count
			}
		}
	}
	return max
}
