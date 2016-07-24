package game

type ReachedMatcher interface {
  Select(board *Board) *MatchedResult
}
