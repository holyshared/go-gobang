package game

type ReachedMatcher interface {
  Matches(board *Board) *MatchedResult
}
