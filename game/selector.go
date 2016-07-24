package game

func NewHorizontalSelector(stone Stone, count int) *HorizontalSelector {
  return &HorizontalSelector { count: count, stone: stone, }
}

func NewVerticalSelector(stone Stone, count int) *VerticalSelector {
  return &VerticalSelector { count: count, stone: stone }
}

func NewTopLeftDiagonalSelector(stone Stone, count int) *TopLeftDiagonalSelector {
  return &TopLeftDiagonalSelector { count: count, stone: stone, }
}

func NewTopRightDiagonalSelector(stone Stone, count int) *TopRightDiagonalSelector {
  return &TopRightDiagonalSelector { count: count, stone: stone, }
}
