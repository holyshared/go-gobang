package game

import (
  "testing"
)

func TestEmptyNeighborCells(t *testing.T) {
  cells := make([]*Cell, 0)
  neighborCells := make([]*Cell, 0)

  cells = append(cells, &Cell { Point { X: 0, Y: 0, }, Black, })
  cells = append(cells, &Cell { Point { X: 1, Y: 0, }, Black, })
  cells = append(cells, &Cell { Point { X: 2, Y: 0, }, Black, })
  neighborCells = append(neighborCells, &Cell { Point { X: 3, Y: 0, }, Black, })

  res := ReachedResult {
    cells: cells,
    neighborCells: neighborCells,
  }

  emptyCells := res.EmptyNeighborCells()

  if len(emptyCells) > 0 {
    t.Errorf("got %v\nwant %v", len(emptyCells), 0)
  }

  cells = cells[0:0]
  neighborCells = neighborCells[0:0]

  cells = append(cells, &Cell { Point { X: 0, Y: 0, }, Black, })
  cells = append(cells, &Cell { Point { X: 1, Y: 0, }, Black, })
  cells = append(cells, &Cell { Point { X: 2, Y: 0, }, Black, })
  neighborCells = append(neighborCells, &Cell { Point { X: 3, Y: 0, }, 0, })

  res = ReachedResult {
    cells: cells,
    neighborCells: neighborCells,
  }

  emptyCells = res.EmptyNeighborCells()

  if len(emptyCells) <= 0 {
    t.Errorf("got %v\nwant %v", len(emptyCells), 1)
  }
}
