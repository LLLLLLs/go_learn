// Time        : 2019/01/24
// Description :

package valid_20

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	fmt.Println(isValid("(){}{}{}()[]({}){()}"))
}
