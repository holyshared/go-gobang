package gobang

import (
  "testing"
)

func TestHasReachedResults(t *testing.T) {
  reachedResults := make([]*ReachedResult, 0)
  reachedResults = append(reachedResults,  &ReachedResult{
    cells: make([]*Cell, 0),
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

func TestHasEmptyNeighborCell(t *testing.T) {
  reachedResults := make([]*ReachedResult, 0)

  emptyNeighborCells := make([]*Cell, 0)
  emptyNeighborCells = append(emptyNeighborCells, NewCell(1, 0, 0) )

  reachedResults = append(reachedResults,  &ReachedResult{
    cells: make([]*Cell, 0),
    emptyNeighborCells: emptyNeighborCells,
  })

  result := &MatchedResult{ results: reachedResults }

  cell := result.SelectEmptyNeighborCell()

  if cell == nil {
    t.Errorf("got %v\nwant neighbor cell", cell)
  }
}
