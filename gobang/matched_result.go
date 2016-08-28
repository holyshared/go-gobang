package gobang

import (
  "math/rand"
)

type MatchedResult struct {
  results []*ReachedResult
}

func (r *MatchedResult) Merge(otherResult *MatchedResult) {
  r.results = append(r.results, otherResult.results...)
}

func (r *MatchedResult) HasResult() bool {
  return len(r.results) >= 1
}

func (r *MatchedResult) HasEmptyNeighborCell() bool {
  has := false

  for _, result := range r.results {
    if !result.HasEmptyNeighborCell() {
      continue
    }
    has = true
    break
  }
  return has
}

func (r *MatchedResult) SelectEmptyNeighborCell() *Cell {
  cells := make([]*Cell, 0)

  for _, result := range r.results {
    cells = append(cells, result.EmptyNeighborCells()...)
  }

  if len(cells) <= 0 {
    return nil
  }

  if len(cells) <= 1 {
    return cells[0]
  }

  index := rand.Intn(len(cells) - 1)

  return cells[index]
}
