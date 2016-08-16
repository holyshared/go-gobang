package gobang

import (
  "math/rand"
)

func NewNpcAI(ctx *NpcAIContext) *NpcAI {
  return &NpcAI {
    ctx: ctx,
  }
}

type GobangAI interface {
  SelectTargetCell() *Cell
}

type NpcAI struct {
  ctx *NpcAIContext
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
//  board := ai.game.CurrentBoard()
//  gamePlayer := ai.game.GamePlayer()

  matcher := NewCellReachedMatcher(ai.ctx.PlayerStone(), ai.ctx.ReachedStoneCount() - 1)
  result := matcher.Matches(ai.ctx.Board())

  if !result.HasEmptyNeighborCell() {
    return nil
  }
  return result.SelectEmptyNeighborCell()
}

func (ai *NpcAI) selectEmptyCell() *Cell {
  board := ai.ctx.Board()
  cells := board.SelectCells(EmptyCell())
  index := rand.Intn(len(cells) - 1)
  return cells[index]
}

func (ai *NpcAI) selectNpcPlayerReachedCell() *Cell {
  var result *MatchedResult

  //board := ai.game.CurrentBoard()
//  npcPlayer := ai.game.NpcPlayer()

  for i := ai.ctx.ReachedStoneCount() - 1; i <= 0; i-- {
    matcher := NewCellReachedMatcher(ai.ctx.NpcPlayerStone(), i)
    result = matcher.Matches(ai.ctx.Board())

    if !result.HasEmptyNeighborCell() {
      continue
    }
    break;
  }

  return result.SelectEmptyNeighborCell()
}
