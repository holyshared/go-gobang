package gobang

func NewNpcArtificialIntelligence() *NpcArtificialIntelligence {
  return &NpcArtificialIntelligence {}
}

type GobangArtificialIntelligence interface {
  SelectTargetCell() *Cell
}

type NpcArtificialIntelligence struct {
  game *GameContext
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

  matcher := NewCellReachedMatcher(gamePlayer.stone, ai.game.ReachedStoneCount() - 1)
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

  for i := ai.game.ReachedStoneCount() - 1; i <= 0; i-- {
    matcher := NewCellReachedMatcher(npcPlayer.stone, i)
    result = matcher.Matches(board)

    if !result.HasEmptyNeighborCell() {
      continue
    }
    break;
  }

  return result.SelectEmptyNeighborCell()
}
