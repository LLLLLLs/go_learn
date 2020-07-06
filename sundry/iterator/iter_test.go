//@author: lls
//@time: 2020/05/29
//@desc:

package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	store := NewStore()
	store.Add(NewItem("apple", 10))
	store.Add(NewItem("orange", 3))
	store.Add(NewItem("mango", 6))
	store.Add(NewItem("watermelon", 2))

	it := store.ItemIter()
	for item := it.Next(); item != nil; item = it.Next() {
		fmt.Println(item.Name(), "price is", item.Price())
	}
	fmt.Println("reset")
	it = store.ItemIter()
	for item := it.Next(); item != nil; item = it.Next() {
		fmt.Println(item.Name(), "price is", item.Price())
	}
}

func TestIsNil(t *testing.T) {
	fmt.Println(getItem() == nil)
}

func getItem() Item {
	return nil
}
