// Time        : 2019/07/30
// Description :

package evaluate_reverse_polish_notation_150

import (
	"fmt"
	"testing"
)

func TestEvaluate(t *testing.T) {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
