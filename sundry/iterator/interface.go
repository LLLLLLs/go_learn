//@author: lls
//@time: 2020/05/29
//@desc:

package iterator

type Store interface {
	Add(item Item)
	ItemIter() Iter
}

type Iter interface {
	Next() Item
}

type Item interface {
	Name() string
	Price() int
}
