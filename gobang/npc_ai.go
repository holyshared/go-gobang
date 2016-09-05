package gobang

import (
	"math/rand"
)

func NewNpcAI(ctx *NpcAIContext) *NpcAI {
	return &NpcAI{
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

	if cell = ai.selectNpcPlayerReachedCell(); cell != nil {
		return cell
	}

	if cell = ai.selectGamePlayerBlockCell(); cell != nil {
		return cell
	}

	if cell = ai.selectNpcPlayerReachChanceCell(); cell != nil {
		return cell
	}

	return ai.selectEmptyCell()
}

func (ai *NpcAI) selectGamePlayerBlockCell() *Cell {
	var cell *Cell

	if cell = ai.selectGamePlayerReachedCell(); cell != nil {
		return cell
	}

	if cell = ai.selectGamePlayerReachChanceCell(); cell != nil {
		return cell
	}

	return nil
}

func (ai *NpcAI) selectGamePlayerReachedCell() *Cell {
	matcher := NewCellReachedMatcher(ai.ctx.PlayerStone(), ai.ctx.ReachedStoneCount()-OneSide)
	result := matcher.Matches(ai.ctx.Board())

	oneSideResult := result.SelectOnly(OneSide)

	if oneSideResult.HasEmptyNeighborCell() {
		return result.SelectEmptyNeighborCell()
	}

	return nil
}

func (ai *NpcAI) selectGamePlayerReachChanceCell() *Cell {
	matcher := NewCellReachedMatcher(ai.ctx.PlayerStone(), ai.ctx.ReachedStoneCount()-BothSides)
	result := matcher.Matches(ai.ctx.Board())

	bothSidesResult := result.SelectOnly(BothSides)

	if bothSidesResult.HasEmptyNeighborCell() {
		return bothSidesResult.SelectEmptyNeighborCell()
	}

	return nil
}

func (ai *NpcAI) selectNpcPlayerReachedCell() *Cell {
	matcher := NewCellReachedMatcher(ai.ctx.NpcPlayerStone(), ai.ctx.ReachedStoneCount()-1)
	result := matcher.Matches(ai.ctx.Board())

	if result.HasEmptyNeighborCell() {
		return result.SelectEmptyNeighborCell()
	}

	return nil
}

func (ai *NpcAI) selectNpcPlayerReachChanceCell() *Cell {
	// 3, 2, 1 ....
	for i := ai.ctx.ReachedStoneCount() - 2; i > 0; i-- {
		matcher := NewCellReachedMatcher(ai.ctx.NpcPlayerStone(), i)
		result := matcher.Matches(ai.ctx.Board())

		// 5 - 3 = 2
		// 5 - 2 = 3
		// 5 - 1 = 4
		remainCellCount := ai.ctx.ReachedStoneCount() - i

		if !result.HaveReachedRemainCell(remainCellCount) {
			continue
		}

		return result.SelectEmptyNeighborCell()
	}

	return nil
}

func (ai *NpcAI) selectEmptyCell() *Cell {
	board := ai.ctx.Board()
	cells := board.SelectCells(EmptyCell())
	index := rand.Intn(len(cells) - 1)

	return cells[index]
}
