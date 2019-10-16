// Time        : 2019/09/25
// Description :

package calc_pi

import (
	"fmt"
	"math/rand"
	"time"
)

// 假设半径=1,点落在圆内相当于 x2 + y2 <= 1;计算1000w个点落在圆内的概率 = [pi*power(r,2)]/power(2r,2) = pi/4
func CalcPiWithRand() {
	rand.Seed(time.Now().Unix())
	inCircle := func(x, y float64) bool {
		if x*x+y*y <= 1 {
			return true
		}
		return false
	}
	var totalCount = 10000000
	var inCircleCount int
	for i := 0; i < totalCount; i++ {
		x, y := rand.Float64(), rand.Float64()
		if inCircle(x, y) {
			inCircleCount++
		}
	}
	fmt.Println(float64(inCircleCount) / float64(totalCount) * 4)
}
