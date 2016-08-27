package gobang

import (
  "testing"
)

func TestHasReachedResults(t *testing.T) {
  reachedResults := make([]*ReachedResult, 0)
  reachedResults = append(reachedResults,  &ReachedResult{
    cells: make([]*Cell, 0),
    neighborCells: make([]*Cell, 0),
    emptyNeighborCells: make([]*Cell, 0),
  })

  result := &MatchedResult{ results: reachedResults }

  if result.HasResult() != true {
    t.Errorf("got %v\nwant %v", result.HasResult(), true)
  }

  if result.HasEmptyNeighborCell() != false {
    t.Errorf("got %v\nwant %v", result.HasEmptyNeighborCell(), false)
  }

  cell := result.SelectEmptyNeighborCell()

  if cell != nil {
    t.Errorf("got %v\nwant %v", cell, nil)
  }
}


func TestHasNotReachedResults(t *testing.T) {
  reachedResults := make([]*ReachedResult, 0)
  result := &MatchedResult{ results: reachedResults }

  if result.HasResult() != false {
    t.Errorf("got %v\nwant %v", result.HasResult(), false)
  }

  if result.HasEmptyNeighborCell() != false {
    t.Errorf("got %v\nwant %v", result.HasEmptyNeighborCell(), false)
  }

  cell := result.SelectEmptyNeighborCell()

  if cell != nil {
    t.Errorf("got %v\nwant %v", cell, nil)
  }
}
