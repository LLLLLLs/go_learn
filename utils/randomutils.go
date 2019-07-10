// Time        : 2019/06/25
// Description :

package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
