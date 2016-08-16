package gobang

func NewGobangPlayer(stone Stone, selector CellSelector) *GobangPlayer {
  return &GobangPlayer {
    stone: stone,
    selector: selector,
  }
}

type Player interface {
  Stone() Stone
  PutStoneTo(cell *Cell) error
}

type GobangPlayer struct {
  stone Stone
  selector CellSelector
}

func (player *GobangPlayer) Stone() Stone {
  return player.stone
}

func (player *GobangPlayer) SelectBoardCell(point *Point) (*Cell, error) {
  return player.selector.SelectCell(point)
}

func (player *GobangPlayer) PutStoneTo(cell *Cell) error {
  if !cell.IsEmpty() {
    return NewAlreadyPlacedError(cell)
  }
  player.stone.PutTo(cell)

  return nil
}
