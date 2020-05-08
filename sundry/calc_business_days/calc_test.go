//@time:2020/04/07
//@desc:

package calc_business_days

import (
	"fmt"
	"testing"
	"time"
)

const (
	day  = 24 * time.Hour
	week = 7 * day
)

func TestCalcBusinessDays(t *testing.T) {
	end := time.Now()
	start := end.Add(-(week*2 + 9*day))
	fmt.Println(CalcBusinessDays(start, end))
}
