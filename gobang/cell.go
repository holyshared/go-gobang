package gobang

func NewCell(x, y int, stone Stone) *Cell {
  point := &Point {
    X: x,
    Y: y,
  }

  return &Cell {
    point,
    stone,
  }
}

type Cell struct {
  *Point `json:"point"`
  Stone Stone `json:"stone"`
}

func (cell *Cell) IsEmpty() bool {
  return cell.Stone != Black && cell.Stone != White
}

func (cell *Cell) Have(stone Stone) bool {
  return cell.Stone == stone
}
