package battleships

type Row struct {
	Cells []Cell
}

type Cell struct {
	State CellState
}

type CellState string

const (
	Empty      CellState = "empty"
	Shipwrek   CellState = "shipwrek"
	HiddenShip CellState = "hidden"
)
