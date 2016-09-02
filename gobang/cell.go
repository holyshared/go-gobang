package gobang

func NewCell(x, y int, stone Stone) *Cell {
  point := &Point {
    x: x,
    y: y,
  }

  return &Cell {
    Point2D: point,
    Stone: stone,
  }
}

type Cell struct {
  Point2D `json:"point"`
  Stone Stone `json:"stone"`
}

func (cell *Cell) IsEmpty() bool {
  return cell.Stone != Black && cell.Stone != White
}

func (cell *Cell) Have(stone Stone) bool {
  return cell.Stone == stone
}

func (cell *Cell) Point() Point2D {
  return cell.Point2D
}
