// Time        : 2019/01/11
// Description :

package factory

import "fmt"

type product interface {
	add()
	sub()
	print()
}

type factory interface {
	create() product
}

type a struct {
	x int
	y int
	r int
}

func (a *a) add() {
	a.r = a.x + a.y
}

func (a *a) sub() {
	a.r = a.x - a.y
}

func (a a) print() {
	fmt.Println(a.r)
}

type aFactory struct {
}

func (af aFactory) create() product {
	return &a{1, 2, 0}
}

type b struct {
	m int
	n int
	r int
}

func (b *b) add() {
	b.r = b.n + b.m
}

func (b *b) sub() {
	b.r = b.n - b.m
}

func (b b) print() {
	fmt.Println(b.r)
}

type bFactory struct {
}

func (bf bFactory) create() product {
	return &b{1, 2, 0}
}
