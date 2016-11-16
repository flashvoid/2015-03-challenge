package battleships

type DefMap struct{}

func (m *DefMap) Shuffle() {
	//
}

func (m *DefMap) GetDimentions() (x, y int) {
	return 0, 0
}

func (m *DefMap) Shoot(x, y int) int {
	return 0
}

func (m *DefMap) Report() []Row {
	return []Row{}
}

func MakeNewMap() Map {
	m := DefMap{}
	m.Shuffle()

	return &m
}
