package game

type GamePlayer struct {
  game *Game
  stone Stone
}

func (p *GamePlayer) PutStone(x, y int) (GameResult, error) {
  board := p.game.CurrentBoard()

  if !board.Have(x, y) {
    return 0, NewCellNotFoundError(NewPoint(x, y))
  }

  if !board.IsCellEmpty(x, y) {
    return 0, NewAlreadyPlacedError(NewPoint(x, y))
  }

  cell := board.Select(x, y)
  p.stone.PutTo(cell)

  matcher := NewCellReachedMatcher(p.stone, 5)
  result := matcher.Matches(board)

  if result.HasResult() {
    return Win, nil
  }

  if board.IsAllFilled() {
    return Draw, nil
  }

  return Skip, nil
}
