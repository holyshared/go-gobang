package gobang

func NewGobangPlayer(stone Stone) *GobangPlayer {
  return &GobangPlayer {
    stone: stone,
  }
}

type Player interface {
  PutStone(x, y int) (PutStoneResult, error)
}

type GobangPlayer struct {
  stone Stone
  game *GameContext
}

func (player *GobangPlayer) JoinTo(game *GameContext) {
  player.game = game
}

func (player *GobangPlayer) PutStone(x, y int) (PutStoneResult, error) {
  board := player.game.CurrentBoard()

  if !board.Have(x, y) {
    return Failed, NewCellNotFoundError(NewPoint(x, y))
  }

  if !board.IsCellEmpty(x, y) {
    return Failed, NewAlreadyPlacedError(NewPoint(x, y))
  }

  cell := board.Select(x, y)
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
