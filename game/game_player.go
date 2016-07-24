package game

import (
  "errors"
)

type GamePlayer struct {
  game *Game
  stone Stone
}

func (p *GamePlayer) PutStone(x, y int) (GameResult, error) {
  board := p.game.CurrentBoard()

  if !board.Have(x, y) {
    return 0, errors.New("You have specified a not exist cell")
  }

  if !board.IsCellEmpty(x, y) {
    return 0, errors.New("Already the stone is placed")
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
