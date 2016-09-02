package gobang

type ReachedSelector struct {
  count int
  stone Stone
  board *Board
  neighbor NeighborDistance
}

func (s ReachedSelector) Select(group *CellGroup) []*ReachedResult {
  builder := &ReachedResultBuilder { count: s.count }
  results := make([]*ReachedResult, 0)

  for _, cell := range group.cells {
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
      builder.AddNeighborCell(p)
    }

    last := builder.LastCell()
    nextPoint := s.neighbor.nextPoint(last)

    if s.board.HaveCell(nextPoint) {
      p := s.board.SelectCell(nextPoint)
      builder.AddNeighborCell(p)
    }

    results = append(results, builder.ReachedResult())
    builder = &ReachedResultBuilder { count: s.count }
  }

  return results;
}
