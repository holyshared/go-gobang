package gobang

type ReachedSelector struct {
	count    int
	stone    Stone
	board    *Board
	neighbor NeighborDistance
}

func (s ReachedSelector) Select(group *CellGroup) []*ReachedResult {
	builder := &ReachedResultBuilder{count: s.count}
	results := make([]*ReachedResult, 0)

	for num, cell := range group.cells {
		if cell.HaveStone(s.stone) == false {
			builder.Clear()
			continue
		}
		builder.AddCell(cell)

		if !builder.IsReached() {
			continue
		}

		first := builder.FirstCell()
		prevPoint := s.neighbor.prevPoint(first)

		if s.board.HaveCell(prevPoint) {
			p := s.board.SelectCell(prevPoint)
			builder.SetFirstNeighborCell(p)
		}

		last := builder.LastCell()
		nextPoint := s.neighbor.nextPoint(last)

		if s.board.HaveCell(nextPoint) {
			p := s.board.SelectCell(nextPoint)
			builder.SetLastNeighborCell(p)
		}

		emptyCount := 0
		firstIndex := num - (s.count - 1)
		lastIndex := firstIndex + s.count

		emptyCount += group.countEmptyCellToFirst(firstIndex)
		emptyCount += group.countEmptyCellToLast(lastIndex)
		builder.SetContinuousEmptyCellCount(emptyCount)

		results = append(results, builder.ReachedResult())
		builder = &ReachedResultBuilder{count: s.count}
	}

	return results
}
