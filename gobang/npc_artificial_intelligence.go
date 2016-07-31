package gobang

func NewNpcArtificialIntelligence(game *Game) *NpcArtificialIntelligence {
  return &NpcArtificialIntelligence {
    game: game,
  }
}

type NpcArtificialIntelligence struct {
  game *Game
}

func (ai *NpcArtificialIntelligence) SelectTargetCell() *Cell {
  cell := ai.selectGamePlayerReachedCell()

  if cell != nil {
    return cell
  }

  return ai.selectNpcPlayerReachedCell()
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

func (ai *NpcArtificialIntelligence) selectNpcPlayerReachedCell() *Cell {
  var result *MatchedResult

  board := ai.game.CurrentBoard()
  npcPlayer := ai.game.NpcPlayer()

  for i := 4; i <= 0; i-- {
    matcher := NewCellReachedMatcher(npcPlayer.stone, i)
    result = matcher.Matches(board)

    if !result.HasEmptyNeighborCell() {
      continue
    }
    break;
  }

  return result.SelectEmptyNeighborCell()
}
