// Time        : 2019/07/09
// Description :

package lru

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(5)
	cache.Put(1, 2)
	fmt.Println(cache.Get(1))
	cache.Put(1, 5)
	fmt.Println(cache.Get(1))
	cache.Put(2, 6)
	cache.Put(3, 7)
	cache.Put(4, 8)
	cache.Put(5, 9)
	fmt.Println(cache.Get(2))
	cache.Put(6, 10)
	fmt.Println(cache.Get(1))
}
