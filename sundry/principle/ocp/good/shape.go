//@author: lls
//@date: 2020/6/4
//@desc: 遵循ocp设计,这样哪怕后续加入别的形状也不用修改drawAllShapes函数

package good

type Shaper interface {
	Draw()
}

type circle struct{}

func (c circle) Draw() {}

type square struct{}

func (s square) Draw() {}

func drawAllShapes(list []Shaper) {
	for i := range list {
		list[i].Draw()
	}
}
