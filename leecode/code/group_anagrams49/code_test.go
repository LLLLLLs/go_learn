// Time        : 2019/07/03
// Description :

package group_anagrams49

import (
	"go_learn/utils"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	res := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	utils.Print2DimensionList(res)
}
