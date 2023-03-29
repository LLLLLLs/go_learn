package module1

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"golearn/pkgtest/fx/work/database"
	"golearn/pkgtest/fx/work/module/agent"
)

type m1 struct {
	db *database.Database
	ms agent.MS
}

func NewM1(db *database.Database) agent.M1 {
	return &m1{db: db}
}

func Init(lc fx.Lifecycle, m agent.M1, ms agent.MS) {
	m.(*m1).ms = ms
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			m.Hello()
			m.(*m1).ms.M2.World()
			return nil
		},
		OnStop: nil,
	})
}

func (m m1) Hello() string {
	str := m.db.Get("hello")
	fmt.Println(str)
	return str
}
