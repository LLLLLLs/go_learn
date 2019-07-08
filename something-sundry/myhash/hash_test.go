// Time        : 2019/01/21
// Description :

package myhash

import (
	"fmt"
	"go_learn/utils"
	"testing"
)

func TestMyHash(t *testing.T) {
	i := 0
	defer func() { fmt.Println(i) }()
	h := newHash()
	for ; i < 100; i++ {
		h.set(i, fmt.Sprintf("%d", i))
	}
}

func BenchmarkMyHash(b *testing.B) {
	h := newHash()
	for i := 0; i < b.N; i++ {
		h.set(i+utils.RandInt(0, 10), fmt.Sprintf("%d", i))
	}
}

func BenchmarkMap(b *testing.B) {
	m := make(map[int]string)
	for i := 0; i < b.N; i++ {
		m[i+utils.RandInt(0, 10)] = fmt.Sprintf("%d", i)
	}
}
