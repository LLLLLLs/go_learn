// Time        : 2019/07/03
// Description :

package group_anagrams49

import (
	"strconv"
	"strings"
)

// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output:
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]

func groupAnagrams(ss []string) [][]string {
	var result = make(map[string][]string)
	for i := range ss {
		key := genKey(ss[i])
		list, ok := result[key]
		if !ok {
			result[key] = []string{ss[i]}
		} else {
			list = append(list, ss[i])
			result[key] = list
		}
	}
	res := make([][]string, 0, len(result))
	for _, r := range result {
		res = append(res, r)
	}
	return res
}

func genKey(s string) string {
	count := make([]int, 26)
	for i := range s {
		count[s[i]-'a']++
	}
	builder := strings.Builder{}
	for i := range count {
		builder.WriteByte('#')
		builder.WriteString(strconv.Itoa(count[i]))
	}
	return builder.String()
}
