package module2

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"golearn/pkgtest/fx/work/database"
	"golearn/pkgtest/fx/work/module/agent"
)

type m2 struct {
	db *database.Database
	ms agent.MS
}

func NewM2(db *database.Database) agent.M2 {
	return &m2{db: db}
}

func Init(lc fx.Lifecycle, m agent.M2, ms agent.MS) {
	m.(*m2).ms = ms
	lc.Append(fx.Hook{
		OnStart: nil,
		OnStop: func(ctx context.Context) error {
			m.World()
			m.(*m2).ms.M1.Hello()
			return nil
		},
	})
}

func (m m2) World() string {
	str := m.db.Get("world")
	fmt.Println(str)
	return str
}
