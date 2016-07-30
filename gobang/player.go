package gobang

type Player interface {
  PutStone(x, y int) (PutStoneResult, error)
}
