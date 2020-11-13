// Time        : 2019/01/10
// Description :

package float

import "math"

func addFloat() {
	const (
		x float64 = 1.1
		z float64 = 1.123
	)
	y := x
	for i := 0; i < 10000000; i++ {
		y *= x
		y /= z
		y += 0.01
		y -= 0.01
	}
}

func addZero() {
	const (
		x float64 = 1.1
		z float64 = 1.123
	)
	y := x
	for i := 0; i < 10000000; i++ {
		y *= x
		y /= z
		y += 0
		y -= 0
	}
}

func float64ToInt64(f64 float64) int64 {
	bigger := f64 >= math.MaxInt64
	if bigger {
		return math.MaxInt64
	}
	return int64(f64)
}
