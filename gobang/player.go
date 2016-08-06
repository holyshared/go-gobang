package gobang

func NewGobangPlayer(stone Stone) *GobangPlayer {
  return &GobangPlayer {
    stone: stone,
  }
}

type Player interface {
  Stone() Stone
  PutStoneTo(cell *Cell) error
}

type GobangPlayer struct {
  stone Stone
  game *GameContext
}

func (player *GobangPlayer) Stone() Stone {
  return player.stone
}

func (player *GobangPlayer) JoinTo(game *GameContext) {
  player.game = game
}

func (player *GobangPlayer) SelectBoardCell(point *Point) (*Cell, error) {
  return player.game.SelectBoardCell(point)
}

func (player *GobangPlayer) PutStoneTo(cell *Cell) error {
  if !cell.IsEmpty() {
    return NewAlreadyPlacedError(cell)
  }
  player.stone.PutTo(cell)

  return nil
}
