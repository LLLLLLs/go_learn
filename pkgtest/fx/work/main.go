package main

import (
	"context"
	"go.uber.org/fx"
	"golearn/pkgtest/fx/work/database"
	"golearn/pkgtest/fx/work/module"
)

func main() {
	provide := []interface{}{database.NewDatabase}
	invoke := make([]interface{}, 0)
	for i := range module.Modules {
		provide = append(provide, module.Modules[i].Provider...)
		invoke = append(invoke, module.Modules[i].Invoker...)
	}
	app := fx.New(
		// fx.Supply(()),
		fx.Provide(provide...),
		fx.Invoke(invoke...),
	)
	if err := app.Start(context.Background()); err != nil {
		panic(err)
	}
	if err := app.Stop(context.Background()); err != nil {
		panic(err)
	}
}
