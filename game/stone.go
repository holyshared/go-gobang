package game

type Stone int

const (
  Black Stone = iota
  White
  None
)

func (stone Stone) Eq(other Stone) bool  {
  return stone == other
}

func (stone Stone) PutTo(cell *Cell) {
  cell.stone = stone
}
