// Time        : 2019/07/25
// Description :

package gas_station_134

import (
	"fmt"
	"testing"
)

func TestGas(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
