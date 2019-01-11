// Time        : 2019/01/10
// Description :

package float

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
