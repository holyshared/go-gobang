package game

type MatchedResult struct {
  results []*ReachedResult
}

func (r *MatchedResult) Merge(otherResult *MatchedResult) {
  r.results = append(r.results, otherResult.results...)
}

func (r *MatchedResult) HasResult() bool {
  return len(r.results) >= 1
}
