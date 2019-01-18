// Time        : 2019/01/11
// Description :

package eng_name

import (
	"fmt"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	s := strings.Replace(str, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, " ", "", -1)
	ss := strings.Split(s, ",")
	nameWithLetter := make([][]string, 0)
	var first = 'A'
	var index = 0
	nameWithLetter = append(nameWithLetter, make([]string, 0))
	for i := range ss {
		sOld := []rune(ss[i])
		var group []string
		if sOld[2] == first {
			group = nameWithLetter[index]
		} else {
			first++
			group = make([]string, 0)
			index++
			nameWithLetter = append(nameWithLetter, group)
		}
		if (i+1)%10 == 0 {
			sOld = append(sOld[:1], sOld[2:len(sOld)-1]...)
		} else {
			sOld = append(sOld[:1], sOld[2:]...)
		}
		sNew := fmt.Sprintf("(%s),", string(sOld))
		if len(group) != 0 && sNew == group[len(group)-1] {
			continue
		}
		group = append(group, sNew)
		nameWithLetter[index] = group
	}
	for i := range nameWithLetter {
		fmt.Println(nameWithLetter[i])
	}
}
