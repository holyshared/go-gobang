package gobang

func NewVerticalCellMatcher(stone Stone, count int) *VerticalCellMatcher {
  return &VerticalCellMatcher { count: count, stone: stone }
}

type VerticalCellMatcher struct {
  count int
  stone Stone
}

func (s *VerticalCellMatcher) Matches(board *Board) *MatchedResult {
  result := MatchedResult {}
  groups := s.scanYAxisCellGroup(board)

  reachedSelector := ReachedSelector {
    stone: s.stone,
    count: s.count,
    board: board,
    neighbor: NewVerticalNeighborDistance(),
  }

  for _, group := range groups {
    results := group.SelectReached(reachedSelector)

    if len(results) <= 0 {
      continue
    }
    result.results = append(result.results, results...)
  }

  return &result
}

func (s *VerticalCellMatcher) scanYAxisCellGroup(board *Board) []*CellGroup {
  endY := board.Height() - 1
  endX := board.Width() - 1
  groups := make([]*CellGroup, 0)

  for x := 0; x <= endX; x++ {
    group := &CellGroup {}

    for y := 0; y <= endY; y++ {
      cell := board.Select(x, y)
      group.cells = append(group.cells, cell)
    }
    groups = append(groups, group)
    group = &CellGroup {}
  }

  return groups
}
