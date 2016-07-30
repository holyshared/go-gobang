package game

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

    if s.board.Have(prevPoint.X, prevPoint.Y) {
      p := s.board.Select(prevPoint.X, prevPoint.Y)
      result.AddNeighborCell(p)
    }

    last := result.Last()
    nextPoint := s.neighbor.nextPoint(last)

    if s.board.Have(nextPoint.X, nextPoint.Y) {
      p := s.board.Select(nextPoint.X, nextPoint.Y)
      result.AddNeighborCell(p)
    }
    results = append(results, result)
    result = &ReachedResult {}
  }

  return results;
}
