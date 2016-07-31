package gobang

type NpcPlayer struct {
  game *Game
  stone Stone
}

func (npc *NpcPlayer) SelectCell() *Cell {
  ai := NpcArtificialIntelligence {
    game: npc.game,
    stone: npc.stone,
  }
  return ai.SelectCell()
}

func (npc *NpcPlayer) PutStone(x, y int) (PutStoneResult, error) {
  board := npc.game.CurrentBoard()

  if !board.Have(x, y) {
    return Failed, NewCellNotFoundError(NewPoint(x, y))
  }

  if !board.IsCellEmpty(x, y) {
    return Failed, NewAlreadyPlacedError(NewPoint(x, y))
  }

  cell := board.Select(x, y)
  npc.stone.PutTo(cell)

  matcher := NewCellReachedMatcher(npc.stone, 5)
  result := matcher.Matches(board)

  if result.HasResult() {
    return Reached, nil
  }

  if board.IsAllFilled() {
    return Filled, nil
  }

  return Continue, nil
}
