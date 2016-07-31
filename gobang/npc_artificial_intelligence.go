package gobang

type NpcArtificialIntelligence struct {
  game *Game
  stone Stone
}

func (ai *NpcArtificialIntelligence) SelectCell() *Cell {
  var result *MatchedResult

  board := ai.game.CurrentBoard()
  cell := ai.selectGamePlayerReachedCell()

  if cell != nil {
    return cell
  }

  for i := 4; i <= 0; i-- {
    matcher := NewCellReachedMatcher(ai.stone, i)
    result = matcher.Matches(board)

    if !result.HasEmptyNeighborCell() {
      continue
    }
    break;
  }

  return result.SelectEmptyNeighborCell()
}

func (ai *NpcArtificialIntelligence) selectGamePlayerReachedCell() *Cell {
  board := ai.game.CurrentBoard()
  gamePlayer := ai.game.GamePlayer()

  matcher := NewCellReachedMatcher(gamePlayer.stone, 4)
  result := matcher.Matches(board)

  if !result.HasEmptyNeighborCell() {
    return nil
  }
  return result.SelectEmptyNeighborCell()
}
