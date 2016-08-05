package gobang

type ReachedSelector struct {
  count int
  stone Stone
  board *Board
  neighbor NeighborDistance
}

func (s ReachedSelector) Select(group *CellGroup) []*ReachedResult {
  result := &ReachedResult {}
  results := make([]*ReachedResult, 0)

  for _, cell := range group.cells {
    if cell.Have(s.stone) == false {
      result.Clear()
      continue
    }
    result.AddCell(cell)

    if !result.IsReached(s.count) {
      continue
    }

    first := result.First()
    prevPoint := s.neighbor.prevPoint(first)

    if s.board.Have(prevPoint) {
      p := s.board.Select(prevPoint)
      result.AddNeighborCell(p)
    }

    last := result.Last()
    nextPoint := s.neighbor.nextPoint(last)

    if s.board.Have(nextPoint) {
      p := s.board.Select(nextPoint)
      result.AddNeighborCell(p)
    }
    results = append(results, result)
    result = &ReachedResult {}
  }

  return results;
}
