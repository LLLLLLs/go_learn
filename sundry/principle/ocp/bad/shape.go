//@author: lls
//@date: 2020/6/4
//@desc: 违反ocp的设计,加入后续加入三角形就需要修改drawAllShapes函数

package bad

type ShapeEnum int

const (
	Circle ShapeEnum = iota
	Square
)

type Shaper interface {
	Shape() ShapeEnum
}

type circle struct {
	Shaper
	radius float64
}

type square struct {
	Shaper
	side float64
}

func drawAllShapes(list []Shaper) {
	for i := range list {
		switch list[i].Shape() {
		case Circle:
			drawCircle(list[i].(circle))
		case Square:
			drawSquare(list[i].(square))
		}
	}
}

func drawCircle(_ circle) {}

func drawSquare(_ square) {}
