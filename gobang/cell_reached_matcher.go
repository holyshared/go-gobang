package gobang

func NewCellReachedMatcher(stone Stone, count int) CellReachedMatcher {
  selectors := []ReachedMatcher {
    NewVerticalCellMatcher(stone, count),
    NewHorizontalCellMatcher(stone, count),
    NewTopLeftDiagonalCellMatcher(stone, count),
    NewTopRightDiagonalCellMatcher(stone, count),
  }

  return CellReachedMatcher { selectors: selectors }
}

type CellReachedMatcher struct {
  selectors []ReachedMatcher
}

func (s *CellReachedMatcher) Matches(board *Board) *MatchedResult {
  result := &MatchedResult {}
  ch := make(chan *MatchedResult, len(s.selectors))

  defer close(ch)

  for _, selector := range s.selectors {
    go func(selector ReachedMatcher, board *Board) {
      ch <- selector.Matches(board)
    }(selector, board)
  }

  for i := len(s.selectors); i > 0; i-- {
    result.Merge(<-ch)
  }

  return result
}
