package gobang

import (
  "testing"
)

func TestEmptyNeighborCells(t *testing.T) {
  cells := make([]*Cell, 0)
  neighborCells := make([]*Cell, 0)

  cells = append(cells, &Cell { NewPoint(0, 0), Black, })
  cells = append(cells, &Cell { NewPoint(1, 0), Black, })
  cells = append(cells, &Cell { NewPoint(2, 0), Black, })
  neighborCells = append(neighborCells, &Cell { NewPoint(3, 0), Black, })

  builder := ReachedResultBuilder {
    cells: cells,
    neighborCells: neighborCells,
  }
  res := builder.ReachedResult()

  emptyCells := res.EmptyNeighborCells()

  if len(emptyCells) > 0 {
    t.Errorf("got %v\nwant %v", len(emptyCells), 0)
  }

  cells = cells[0:0]
  neighborCells = neighborCells[0:0]

  cells = append(cells, &Cell { NewPoint(0, 0), Black, })
  cells = append(cells, &Cell { NewPoint(1, 0), Black, })
  cells = append(cells, &Cell { NewPoint(2, 0), Black, })
  neighborCells = append(neighborCells, &Cell { NewPoint(3, 0), 0, })

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
