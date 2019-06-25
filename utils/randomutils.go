// Time        : 2019/06/25
// Description :

package utils

import "math/rand"

func RandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
