package gobang

func NewGobangPlayer(stone Stone) *GobangPlayer {
  return &GobangPlayer {
    stone: stone,
  }
}

type Player interface {
  PutStoneTo(cell *Cell) (PutStoneResult, error)
}

type GobangPlayer struct {
  stone Stone
  game *GameContext
}

func (player *GobangPlayer) JoinTo(game *GameContext) {
  player.game = game
}

func (player *GobangPlayer) SelectBoardCell(point *Point) (*Cell, error) {
  return player.game.SelectBoardCell(point)
}

func (player *GobangPlayer) PutStoneTo(cell *Cell) (PutStoneResult, error) {
  board := player.game.CurrentBoard()

  if !cell.IsEmpty() {
    return Failed, NewAlreadyPlacedError(cell)
  }
  player.stone.PutTo(cell)

  matcher := NewCellReachedMatcher(player.stone, player.game.ReachedStoneCount())
  result := matcher.Matches(board)

  if result.HasResult() {
    return Reached, nil
  }

  if board.IsAllFilled() {
    return Filled, nil
  }

  return Continue, nil
}
