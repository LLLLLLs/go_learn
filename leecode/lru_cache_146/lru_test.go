// Time        : 2019/07/30
// Description :

package lru_cache_146

import (
	"fmt"
	"testing"
)

//["LRUCache","put","put","get","put","get","put","get","get","get"]
//[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
func TestLruCache(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
}

//["LRUCache","put","get","put","get","get"]
//[[1],[2,1],[2],[3,2],[2],[3]]
func TestLruCache2(t *testing.T) {
	lru := Constructor(1)
	lru.Put(2, 1)
	fmt.Println(lru.Get(2))
	lru.Put(3, 2)
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(3))
}

//["LRUCache","put","put","get","put","put","get"]
//[[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
func TestLruCache3(t *testing.T) {
	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(2))
	lru.Put(1, 1)
	lru.Put(4, 1)
	fmt.Println(lru.Get(2))
}
