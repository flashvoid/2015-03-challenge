package battleships

type DefPlayer struct {
	currentScore int
}

func (p *DefPlayer) AddScore(x int) {
	p.currentScore += x
}

func (p *DefPlayer) GetScore() int {
	return p.currentScore
}

func MakeNewPlayer() Player {
	p := DefPlayer{}
	return &p
}
