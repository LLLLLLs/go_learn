// Time        : 2019/07/04
// Description :

package insert_interval_57

import (
	"fmt"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	output := insert(intervals, newInterval)
	fmt.Println(output)
	intervals = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval = []int{4, 8}
	output = insert(intervals, newInterval)
	fmt.Println(output)
	intervals = [][]int{{2, 6}, {7, 9}}
	newInterval = []int{15, 18}
	output = insert(intervals, newInterval)
	fmt.Println(output)
}
