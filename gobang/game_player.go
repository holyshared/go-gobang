package gobang

type GamePlayer struct {
  game *GameContext
  stone Stone
}

func (p *GamePlayer) PutStone(x, y int) (PutStoneResult, error) {
  board := p.game.CurrentBoard()

  if !board.Have(x, y) {
    return Failed, NewCellNotFoundError(NewPoint(x, y))
  }

  if !board.IsCellEmpty(x, y) {
    return Failed, NewAlreadyPlacedError(NewPoint(x, y))
  }

  cell := board.Select(x, y)
  p.stone.PutTo(cell)

  matcher := NewCellReachedMatcher(p.stone, 5)
  result := matcher.Matches(board)

  if result.HasResult() {
    return Reached, nil
  }

  if board.IsAllFilled() {
    return Filled, nil
  }

  return Continue, nil
}
