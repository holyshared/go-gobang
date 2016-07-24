package game

/**
 * | | | | | | | |B| |B|
 * | | | | | | |B| |B| |
 * | | | | | |B| |B| | |
 * | | | | |B| |B| | | |
 * | | | |B| |B| | | | |
 * | | | | |B| | | | | |
 * | | | |B| | | | | | |
 * | | |B| | | | | | | |
 * | |B| | | | | | | | |
 * |B| | | | | | | | | |
 */
type TopRightDiagonalSelector struct {
  count int
  stone Stone
}

func (s *TopRightDiagonalSelector) Matches(board *Board) *MatchedResult {
  result := &MatchedResult {}
  groups := s.scanAllCellGroup(board)

  reachedSelector := ReachedSelector {
    stone: s.stone,
    count: s.count,
    board: board,
    neighbor: NewTopRightNeighborDistance(),
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


func (s *TopRightDiagonalSelector) scanAllCellGroup(board *Board) []*CellGroup {
  groups := make([]*CellGroup, 0)
  groups = append(groups, s.scanXAxisCellGroup(board)...)
  groups = append(groups, s.scanYAxisCellGroup(board)...)
  return groups
}

/**
 * | | | | |B|B|B|B|B|B|
 * | | | |B|B|B|B|B|B| |
 * | | |B|B|B|B|B|B| | |
 * | |B|B|B|B|B|B| | | |
 * |B|B|B|B|B|B| | | | |
 * |B|B|B|B|B| | | | | |
 * |B|B|B|B| | | | | | |
 * |B|B|B| | | | | | | |
 * |B|B| | | | | | | | |
 * |B| | | | | | | | | |
 */
func (s *TopRightDiagonalSelector) scanXAxisCellGroup(board *Board) []*CellGroup {
  endX := board.Width() - 1
  groups := make([]*CellGroup, 0)

  for startX := s.count - 1; startX <= endX; startX++ {
    y := 0
    group := &CellGroup {}
    for x := startX; x >= 0; x-- {
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
 * | | | | | | | | | |B|
 * | | | | | | | | |B|B|
 * | | | | | | | |B|B|B|
 * | | | | | | |B|B|B|B|
 * | | | | | |B|B|B|B|B|
 * | | | | |B|B|B|B|B| |
 * | | | |B|B|B|B|B| | |
 * | | |B|B|B|B|B| | | |
 * | |B|B|B|B|B| | | | |
 */
func (s *TopRightDiagonalSelector) scanYAxisCellGroup(board *Board) []*CellGroup {
  maxY := board.Height() - 1
  endY := board.Height() - s.count
  groups := make([]*CellGroup, 0)

  for startY := 1; startY <= endY; startY++ {
    x := board.Width() - 1
    group := &CellGroup {}

    for y := startY; y <= maxY; y++ {
      cell := board.Select(x, y)
      group.cells = append(group.cells, cell)
      x--
    }
    groups = append(groups, group)
    group = &CellGroup {}
  }

  return groups
}
