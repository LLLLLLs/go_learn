// Time        : 2019/07/09
// Description :

package plusone_66

import (
	"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {
	input := []int{1, 2, 3, 4}
	output := plusOne(input)
	fmt.Println(output)

	input = []int{9, 9, 9, 9}
	output = plusOne(input)
	fmt.Println(output)
}
