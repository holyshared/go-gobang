package game

type MatchedResult struct {
  results []*ReachedResult
}

func (r *MatchedResult) Merge(otherResult *MatchedResult) {
  r.results = append(r.results, otherResult.results...)
}
