package gobang

type ReachedMatcher interface {
  Matches(board *Board) *MatchedResult
}
