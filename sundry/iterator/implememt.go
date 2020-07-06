//@author: lls
//@time: 2020/05/29
//@desc:

package iterator

type store struct {
	items []Item
}

func NewStore() Store {
	return &store{
		items: make([]Item, 0),
	}
}

func (s *store) Add(item Item) {
	s.items = append(s.items, item)
}

func (s store) ItemIter() Iter {
	return &iter{
		index: -1,
		store: s,
	}
}

type iter struct {
	index int
	store store
}

func (i *iter) Next() Item {
	if i.index+1 >= len(i.store.items) {
		return nil
	}
	i.index++
	return i.store.items[i.index]
}

type item struct {
	name  string
	price int
}

func NewItem(name string, price int) Item {
	return item{
		name:  name,
		price: price,
	}
}

func (i item) Name() string {
	return i.name
}

func (i item) Price() int {
	return i.price
}
