package game

type Player interface {
  PutStone(x, y int) (PutStoneResult, error)
}
