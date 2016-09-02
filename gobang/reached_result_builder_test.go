package gobang

import (
  "testing"
)

func TestEmptyNeighborCells(t *testing.T) {
  cells := []*Cell {
    NewCell(0, 0, Black),
    NewCell(1, 0, Black),
    NewCell(2, 0, Black),
  }
  neighborCells := []*Cell {
    NewCell(3, 0, Black),
  }

  builder := ReachedResultBuilder {
    cells: cells,
    neighborCells: neighborCells,
  }
  res := builder.ReachedResult()

  emptyCells := res.EmptyNeighborCells()

  if len(emptyCells) > 0 {
    t.Errorf("got %v\nwant %v", len(emptyCells), 0)
  }

  cells = []*Cell {
    NewCell(0, 0, Black),
    NewCell(1, 0, Black),
    NewCell(2, 0, Black),
  }
  neighborCells = []*Cell {
    NewCell(3, 0, 0),
  }

  builder = ReachedResultBuilder {
    cells: cells,
    neighborCells: neighborCells,
  }

  res = builder.ReachedResult()
  emptyCells = res.EmptyNeighborCells()

  if len(emptyCells) <= 0 {
    t.Errorf("got %v\nwant %v", len(emptyCells), 1)
  }
}
