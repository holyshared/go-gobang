package game

func NewCellReachedMatcher(stone Stone, count int) CellReachedMatcher {
  selectors := make([]ReachedMatcher, 0)
  selectors = append(selectors, NewVerticalSelector(stone, count))
  selectors = append(selectors, NewHorizontalSelector(stone, count))
  selectors = append(selectors, NewTopLeftDiagonalSelector(stone, count))
  selectors = append(selectors, NewTopRightDiagonalSelector(stone, count))

  return CellReachedMatcher { selectors: selectors }
}

type CellReachedMatcher struct {
  selectors []ReachedMatcher
}

func (s *CellReachedMatcher) Select(board *Board) *MatchedResult {
  result := &MatchedResult {}

  for _, selector := range s.selectors {
    result.Merge(selector.Select(board))
  }

  return result
}
