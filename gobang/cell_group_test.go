package gobang

import (
  "testing"
)

func TestEmptyNeighborCellCount(t *testing.T) {
  var actual int

  cells := []*Cell {
    NewCell(0, 0, 0),
    NewCell(1, 0, Black),
    NewCell(2, 0, 0),
  }
  group := &CellGroup {
    cells: cells,
  }
  actual = group.countEmptyCellToFirst(0)

  if actual != 1 {
    t.Errorf("got %v\nwant %v", actual, 1)
  }

  actual = group.countEmptyCellToLast(2)

  if actual != 1 {
    t.Errorf("got %v\nwant %v", actual, 1)
  }
}
