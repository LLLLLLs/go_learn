//@author: lls
//@time: 2020/08/27
//@desc:

package compare_version_165

import (
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i := 0; i < max(len(v1), len(v2)); i++ {
		subV1 := subV(v1, i)
		subV2 := subV(v2, i)
		if subV1 == subV2 {
			continue
		}
		if bigger(subV1, subV2) {
			return 1
		}
		return -1
	}
	return 0
}
func subV(v []string, i int) string {
	if i <= len(v)-1 {
		return strings.TrimLeft(v[i], "0")
	}
	return ""
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bigger(s1, s2 string) bool {
	return len(s1) > len(s2) || len(s1) == len(s2) && s1 > s2
}
