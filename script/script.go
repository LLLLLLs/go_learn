package script

import (
	"fmt"
	"sync"
)

type A struct {
	B *B
}

type B struct {
	MM sync.Map
}

func (b *B) Func() {
	if b == nil {
		return
	}
	if v, ok := b.MM.Load("func"); ok {
		fmt.Println("func", v)
	}
}

func (b *B) Empty() {
	fmt.Println(b == nil)
}
