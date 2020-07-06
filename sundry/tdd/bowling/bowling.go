//@author: lls
//@date: 2020/6/3
//@desc: 保龄球比赛计分

package bowling

type Game struct {
	score             Scorer
	currentFrame      int
	firstThrowInFrame bool
}

func NewGame() Game {
	return Game{
		score:             Scorer{},
		currentFrame:      1,
		firstThrowInFrame: true,
	}
}

func (g Game) Score() int {
	return g.ScoreForFrame(g.currentFrame)
}

func (g *Game) Add(pins int) {
	g.score.Add(pins)
	g.adjustCurrentFrame(pins)
}

func (g *Game) adjustCurrentFrame(pins int) {
	if !g.firstThrowInFrame || pins == 10 {
		g.advanceFrame()
	} else {
		g.firstThrowInFrame = false
	}
}

func (g *Game) advanceFrame() {
	g.firstThrowInFrame = true
	g.currentFrame++
	if g.currentFrame > 10 {
		g.currentFrame = 10
	}
}

func (g *Game) ScoreForFrame(frame int) int {
	return g.score.ScoreForFrame(frame)
}

type Scorer struct {
	ball         int
	throws       [21]int
	currentThrow int
}

func (s *Scorer) Add(pins int) {
	s.throws[s.currentThrow] = pins
	s.currentThrow++
}

func (s *Scorer) ScoreForFrame(frame int) int {
	score := 0
	s.ball = 0
	for currentFrame := 0; currentFrame < frame; currentFrame++ {
		score += s.scoreForCurrentFrame()
	}
	return score
}

func (s *Scorer) scoreForCurrentFrame() int {
	var score int
	if s.strike() {
		score += 10 + s.nextTwoBallsForStrike()
		s.ball++
	} else if s.spare() {
		score += 10 + s.nextBallForSpare()
		s.ball += 2
	} else {
		score += s.twoBallsInFrame()
		s.ball += 2
	}
	return score
}

func (s Scorer) strike() bool {
	return s.throws[s.ball] == 10
}

func (s Scorer) spare() bool {
	return s.throws[s.ball]+s.throws[s.ball+1] == 10
}

func (s Scorer) twoBallsInFrame() int {
	return s.throws[s.ball] + s.throws[s.ball+1]
}

func (s Scorer) nextBallForSpare() int {
	return s.throws[s.ball+2]
}

func (s Scorer) nextTwoBallsForStrike() int {
	return s.throws[s.ball+1] + s.throws[s.ball+2]
}
