//@author: lls
//@time: 2020/05/06
//@desc:

package private

type A struct {
	a int
	b string
}

func NewA(a int, b string) A {
	return A{
		a: a,
		b: b,
	}
}
