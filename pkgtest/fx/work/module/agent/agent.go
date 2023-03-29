package agent

import "go.uber.org/fx"

type MS struct {
	fx.In
	M1 M1
	M2 M2
}

type M1 interface {
	Hello() string
}

type M2 interface {
	World() string
}
