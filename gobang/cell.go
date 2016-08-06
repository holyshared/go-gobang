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
  *Point
  stone Stone
}

func (cell *Cell) IsEmpty() bool {
  return cell.stone != Black && cell.stone != White
}

func (cell *Cell) Have(stone Stone) bool {
  return cell.stone == stone
}
