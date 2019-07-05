// Time        : 2019/07/03
// Description :

package count_II_52

import "strings"

func totalNQueens(n int) int {
	count := 0
	recursive([]string{}, n, &count)
	return count
}

func recursive(mid []string, n int, count *int) {
	if len(mid) == n {
		*count++
		return
	}
	for j := 0; j < n; j++ {
		hasQueen := false
		for i := range mid {
			if mid[i][j] == 'Q' {
				hasQueen = true
				break
			}
		}
		for i := 0; !hasQueen && i < n; i++ {
			skewI := len(mid) - abs(i-j)
			if 0 <= skewI && skewI < len(mid) && mid[skewI][i] == 'Q' {
				hasQueen = true
				break
			}
		}
		if !hasQueen {
			recursive(append(append([]string{}, mid...), genCol(j, n)), n, count)
		}
	}
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func genCol(index, n int) string {
	builder := strings.Builder{}
	for i := 0; i < n; i++ {
		if i == index {
			builder.WriteByte('Q')
		} else {
			builder.WriteByte('.')
		}
	}
	return builder.String()
}
