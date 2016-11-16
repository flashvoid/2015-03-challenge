package battleships

type Map interface {
	// GetDimentions reports size of the battlefield.
	GetDimentions(x, y int)

	// Shuffle wipes map state and drops new set of ships on a map randomly.
	Shuffle()

	// Shoot checks if there is a ship on given coordinates and reports hit/miss.
	Shoot(x, y int) bool

	// Reports a state od every cell empty/shipwrek.
	Report() []Row
}

type Player interface {
	// AddScore
	AddScore(int)

	GetScore() int
}

type CnC interface {
	// Init
	// Run starts reading commands from stdin and writing reponses to stdout.
	// * Upon receiving `start` call Map.Shuffle()
	// * Upon receiving `shoot x y` call Map.Shoot(x,y)
	// * * if hit update player score.
	// * Upon receiving `report` call Map.Report() and draw it somehow.
	Run()
}
