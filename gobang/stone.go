package gobang

type Stone int

const (
  Black Stone = iota + 1
  White
)

func (stone Stone) Eq(other Stone) bool  {
  return stone == other
}

func (stone Stone) PutTo(cell *Cell) {
  cell.stone = stone
}

func (stone Stone) String() string {
  if (stone == Black) {
    return "B"
  } else {
    return "W"
  }
}
