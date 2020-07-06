//@author: lls
//@date: 2020/6/3
//@desc:

package bowling_test

import (
	"github.com/stretchr/testify/assert"
	. "golearn/sundry/tdd/bowling"
	"testing"
)

func TestGame(t *testing.T) {
	ast := assert.New(t)

	t.Run("two throws no mark", withNewGame(func(g Game, t *testing.T) {
		g.Add(4)
		g.Add(5)
		ast.Equal(9, g.Score())
	}))

	t.Run("four throws no mark", withNewGame(func(g Game, t *testing.T) {
		g.Add(5)
		g.Add(4)
		g.Add(7)
		g.Add(2)
		ast.Equal(18, g.Score(), "总得分：18")
		ast.Equal(9, g.ScoreForFrame(1), "第一轮：9")
		ast.Equal(18, g.ScoreForFrame(2), "第二轮：9")
	}))

	t.Run("simple spare", withNewGame(func(g Game, t *testing.T) {
		g.Add(3)
		g.Add(7)
		g.Add(3)
		ast.Equal(13, g.ScoreForFrame(1), "第一轮补中：13分")
	}))

	t.Run("simple frame after spare", withNewGame(func(g Game, t *testing.T) {
		g.Add(3)
		g.Add(7)
		g.Add(3)
		g.Add(2)
		ast.Equal(13, g.ScoreForFrame(1))
		ast.Equal(18, g.ScoreForFrame(2))
		ast.Equal(18, g.Score())
	}))

	t.Run("simple strike", withNewGame(func(g Game, t *testing.T) {
		g.Add(10)
		g.Add(3)
		g.Add(6)
		ast.Equal(19, g.ScoreForFrame(1))
		ast.Equal(28, g.Score())
	}))

	t.Run("perfect game", withNewGame(func(g Game, t *testing.T) {
		// 完美比赛，投12次10分，因为第十轮得10分需要多投2次
		for i := 0; i < 12; i++ {
			g.Add(10)
		}
		ast.Equal(300, g.Score())
	}))

	t.Run("end of throw array", withNewGame(func(g Game, t *testing.T) {
		for i := 0; i < 9; i++ {
			g.Add(0)
			g.Add(0)
		}
		g.Add(2)
		g.Add(8)
		g.Add(10)
		ast.Equal(20, g.Score())
	}))

	t.Run("sample game", withNewGame(func(g Game, t *testing.T) {
		g.Add(1)
		g.Add(4)
		g.Add(4)
		g.Add(5)
		g.Add(6)
		g.Add(4)
		g.Add(5)
		g.Add(5)
		g.Add(10)
		g.Add(0)
		g.Add(1)
		g.Add(7)
		g.Add(3)
		g.Add(6)
		g.Add(4)
		g.Add(10)
		g.Add(2)
		g.Add(8)
		g.Add(6)
		ast.Equal(133, g.Score())
	}))

	t.Run("heart break", withNewGame(func(g Game, t *testing.T) {
		for i := 0; i < 11; i++ {
			g.Add(10)
		}
		g.Add(9)
		ast.Equal(299, g.Score())
	}))

	t.Run("tenth frame spare", withNewGame(func(g Game, t *testing.T) {
		for i := 0; i < 9; i++ {
			g.Add(10)
		}
		g.Add(9)
		g.Add(1)
		g.Add(1)
		ast.Equal(270, g.Score())
	}))
}

func withNewGame(f func(g Game, t *testing.T)) func(t *testing.T) {
	g := NewGame()
	return func(t *testing.T) {
		f(g, t)
	}
}
