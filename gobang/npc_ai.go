package gobang

import (
  "math/rand"
)

func NewNpcAI(game *GameContext) *NpcAI {
  return &NpcAI {
    game: game,
  }
}

type GobangAI interface {
  SelectTargetCell() *Cell
}

type NpcAI struct {
  game *GameContext
}

func (ai *NpcAI) SelectTargetCell() *Cell {
  var cell *Cell

  cell = ai.selectGamePlayerReachedCell()

  if cell != nil {
    return cell
  }

  cell = ai.selectNpcPlayerReachedCell()

  if cell != nil {
    return cell
  }

  return ai.selectEmptyCell()
}

func (ai *NpcAI) selectGamePlayerReachedCell() *Cell {
  board := ai.game.CurrentBoard()
  gamePlayer := ai.game.GamePlayer()

  matcher := NewCellReachedMatcher(gamePlayer.stone, ai.game.ReachedStoneCount() - 1)
  result := matcher.Matches(board)

  if !result.HasEmptyNeighborCell() {
    return nil
  }
  return result.SelectEmptyNeighborCell()
}

func (ai *NpcAI) selectEmptyCell() *Cell {
  board := ai.game.CurrentBoard()
  cells := board.SelectCells(EmptyCell())
  index := rand.Intn(len(cells) - 1)
  return cells[index]
}

func (ai *NpcAI) selectNpcPlayerReachedCell() *Cell {
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
