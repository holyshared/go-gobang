package gobang

func NewGobangPlayer(stone Stone) *GobangPlayer {
  return &GobangPlayer {
    stone: stone,
  }
}

type Player interface {
  PutStone(point *Point) (PutStoneResult, error)
}

type GobangPlayer struct {
  stone Stone
  game *GameContext
}

func (player *GobangPlayer) JoinTo(game *GameContext) {
  player.game = game
}

func (player *GobangPlayer) PutStone(point *Point) (PutStoneResult, error) {
  board := player.game.CurrentBoard()

  if !board.Have(point) {
    return Failed, NewCellNotFoundError(point)
  }

  if !board.IsCellEmpty(point) {
    return Failed, NewAlreadyPlacedError(point)
  }
  player.putStoneTo(point)

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

func (player *GobangPlayer) putStoneTo(point *Point) {
  board := player.game.CurrentBoard()

  cell := board.Select(point)
  player.stone.PutTo(cell)
}
