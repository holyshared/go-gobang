package gobang

func NewHorizontalCellMatcher(stone Stone, count int) *HorizontalCellMatcher {
  return &HorizontalCellMatcher { count: count, stone: stone, }
}

type HorizontalCellMatcher struct {
  count int
  stone Stone
}

func (s *HorizontalCellMatcher) Matches(board *Board) *MatchedResult {
  result := MatchedResult {}
  groups := s.scanXAxisCellGroup(board)

  reachedSelector := ReachedSelector {
    stone: s.stone,
    count: s.count,
    board: board,
    neighbor: NewHorizontalNeighborDistance(),
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


func (s *HorizontalCellMatcher) scanXAxisCellGroup(board *Board) []*CellGroup {
  endY := board.Height() - 1
  endX := board.Width() - 1
  groups := make([]*CellGroup, 0)

  for y := 0; y <= endY; y++ {
    group := &CellGroup {}

    for x := 0; x <= endX; x++ {
      cell := board.Select(NewPoint(x, y))
      group.cells = append(group.cells, cell)
    }
    groups = append(groups, group)
    group = &CellGroup {}
  }

  return groups
}
