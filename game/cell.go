package game

type Cell struct {
  stone Stone
  point Point
}

func (cell *Cell) IsEmpty() bool {
  return cell.stone != Black && cell.stone != White
}

func (cell *Cell) Have(stone Stone) bool {
  return cell.stone == stone
}
