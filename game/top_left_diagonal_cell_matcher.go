package game

/**
 * |B| | | | | | | | | |
 * | |B| | | | | | | | |
 * | | |B| | | | | | | |
 * | | | |B| | | | | | |
 * | | | | |B| | | | | |
 * | | | |B| | | | | | |
 * | | | | |B| | | | | |
 * | | | | | |B| | | | |
 * | | | | | | |B| | | |
 * | | | | | | | |B| | |
 */
func NewTopLeftDiagonalCellMatcher(stone Stone, count int) *TopLeftDiagonalCellMatcher {
  return &TopLeftDiagonalCellMatcher { count: count, stone: stone, }
}

type TopLeftDiagonalCellMatcher struct {
  count int
  stone Stone
}

func (s *TopLeftDiagonalCellMatcher) Matches(board *Board) *MatchedResult {
  result := &MatchedResult {}
  groups := s.scanAllCellGroup(board)

  reachedSelector := ReachedSelector {
    stone: s.stone,
    count: s.count,
    board: board,
    neighbor: NewTopLeftNeighborDistance(),
  }

  for _, group := range groups {
    results := group.SelectReached(reachedSelector)

    if len(results) <= 0 {
      continue
    }
    result.results = append(result.results, results...)
  }

  return result
}

func (s *TopLeftDiagonalCellMatcher) scanAllCellGroup(board *Board) []*CellGroup {
  groups := make([]*CellGroup, 0)
  groups = append(groups, s.scanXAxisCellGroup(board)...)
  groups = append(groups, s.scanYAxisCellGroup(board)...)
  return groups
}

/**
 * |B|B|B|B|B|B| | | | |
 * | |B|B|B|B|B|B| | | |
 * | | |B|B|B|B|B|B| | |
 * | | | |B|B|B|B|B|B| |
 * | | | | |B|B|B|B|B|B|
 * | | | | | |B|B|B|B|B|
 * | | | | | | |B|B|B|B|
 * | | | | | | | |B|B|B|
 * | | | | | | | | |B|B|
 * | | | | | | | | | |B|
 */
func (s *TopLeftDiagonalCellMatcher) scanXAxisCellGroup(board *Board) []*CellGroup {
  maxX := board.Width() - 1
  endX := board.Width() - s.count
  groups := make([]*CellGroup, 0)

  for startX := 0; startX <= endX; startX++ {
    y := 0
    endY := (board.Height() - startX)
    group := &CellGroup {}

    for x := startX; x <= maxX; x++ {
      if y > endY {
        break
      }
      cell := board.Select(x, y)
      group.cells = append(group.cells, cell)
      y++
    }
    groups = append(groups, group)
    group = &CellGroup {}
  }

  return groups
}

/**
 * | | | | | | | | | | |
 * |B| | | | | | | | | |
 * |B|B| | | | | | | | |
 * |B|B|B| | | | | | | |
 * |B|B|B|B| | | | | | |
 * |B|B|B|B|B| | | | | |
 * | |B|B|B|B|B| | | | |
 * | | |B|B|B|B|B| | | |
 * | | | |B|B|B|B|B| | |
 * | | | | |B|B|B|B|B| |
 */
func (s *TopLeftDiagonalCellMatcher) scanYAxisCellGroup(board *Board) []*CellGroup {
  maxY := board.Height() - 1
  endY := board.Height() - s.count
  groups := make([]*CellGroup, 0)

  for startY := 1; startY <= endY; startY++ {
    x := 0
    endX := (board.Width() - startY)
    group := &CellGroup {}

    for y := startY; y <= maxY; y++ {
      if x > endX {
        break
      }
      cell := board.Select(x, y)
      group.cells = append(group.cells, cell)
      x++
    }
    groups = append(groups, group)
    group = &CellGroup {}
  }
  return groups
}
