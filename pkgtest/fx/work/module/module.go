package module

import (
	"golearn/pkgtest/fx/work/module/module1"
	"golearn/pkgtest/fx/work/module/module2"
)

type Module struct {
	Provider []interface{}
	Invoker  []interface{}
}

var Modules = []Module{
	{
		Provider: []interface{}{module1.NewM1},
		Invoker:  []interface{}{module1.Init},
	},
	{
		Provider: []interface{}{module2.NewM2},
		Invoker:  []interface{}{module2.Init},
	},
}
