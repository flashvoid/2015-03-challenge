package battleships

type Map interface {
  // GetDimentions reports size of the battlefield.
  GetDimentions(x, y int)
  
  // Shuffle wipes map state and drops new set of ships on a map randomly.
  Shuffle()
  
  // Shoot checks if there is a ship on given coordinates and reports hit/miss.
  Shoot(x, y int) bool
  
  // Reports a state od every cell empty/shipwrek.
  Report() []Rows
}
