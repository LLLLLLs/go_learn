package main

import (
	"fmt"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(NewA, NewB),
		fx.Invoke(testA),
	)
	app.Run()
}

func testA(a IA) {
	a.A()
}

type IA interface {
	A()
}

type IB interface {
	B()
}

type A struct {
	B IB
}

func NewA() IA {
	return &A{}
}

func (a *A) Init(b IB) {
	a.B = b
}

func (a *A) A() {
	a.B.B()
}

type B struct {
	A   IA
	str string
}

func NewB() IB {
	return &B{}
}

func (b *B) Init(a IA) {
	b.A = a
	b.str = "hello"
}

func (b *B) B() {
	fmt.Println(b.str)
}
