package battleships

type DefGame struct {
	turn   int
	player Player
	m      Map
}

func MakeNewGame() *DefGame {
	g := DefGame{}
	g.player = MakeNewPlayer()
	g.m = MakeNewMap()

	return &g
}
