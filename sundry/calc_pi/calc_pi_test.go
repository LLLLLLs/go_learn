// Time        : 2019/09/25
// Description :

package calc_pi

import (
	"fmt"
	"testing"
	"time"
)

func TestCalcPiWithRand(t *testing.T) {
	CalcPiWithRand()
}

func TestTime(t *testing.T) {
	fmt.Println(time.Unix(1588054103, 0))
	fmt.Println(time.Unix(1588054163, 0))
	fmt.Println(time.Unix(1588054180, 0))

}
