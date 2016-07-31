package gobang

type NpcPlayer struct {
  game *GameContext
  stone Stone
}

func (npc *NpcPlayer) SelectTargetCell() *Cell {
  ai := NpcArtificialIntelligence {
    game: npc.game,
  }
  return ai.SelectTargetCell()
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

  matcher := NewCellReachedMatcher(npc.stone, npc.game.ReachedStoneCount())
  result := matcher.Matches(board)

  if result.HasResult() {
    return Reached, nil
  }

  if board.IsAllFilled() {
    return Filled, nil
  }

  return Continue, nil
}
