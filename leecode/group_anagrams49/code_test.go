// Time        : 2019/07/03
// Description :

package group_anagrams49

import (
	"golearn/util"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	res := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	util.Print2DimensionList(res)
}
