//@author: lls
//@time: 2020/04/10
//@desc:

package gob

type IA interface {
	Hello()
}

type A struct {
	Y int
}

func (a A) Hello() {}

type B struct {
	AA IA
	X  int
}
