package game

func NewHorizontalCellMatcher(stone Stone, count int) *HorizontalCellMatcher {
  return &HorizontalCellMatcher { count: count, stone: stone, }
}

func NewVerticalCellMatcher(stone Stone, count int) *VerticalCellMatcher {
  return &VerticalCellMatcher { count: count, stone: stone }
}

func NewTopLeftDiagonalCellMatcher(stone Stone, count int) *TopLeftDiagonalCellMatcher {
  return &TopLeftDiagonalCellMatcher { count: count, stone: stone, }
}

func NewTopRightDiagonalCellMatcher(stone Stone, count int) *TopRightDiagonalCellMatcher {
  return &TopRightDiagonalCellMatcher { count: count, stone: stone, }
}
