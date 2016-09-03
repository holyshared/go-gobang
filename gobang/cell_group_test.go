package gobang

import (
  "testing"
)

func TestEmptyNeighborCellCount(t *testing.T) {
  var actual int

  cells := []*Cell {
    NewCell(0, 0, 0),
    NewCell(1, 0, 0),
    NewCell(2, 0, Black),
    NewCell(3, 0, 0),
    NewCell(4, 0, Black),
    NewCell(5, 0, 0),
    NewCell(6, 0, Black),
    NewCell(7, 0, 0),
  }
  group := &CellGroup {
    cells: cells,
  }
  actual = group.countEmptyCellToFirst(3)

  if actual != 1 {
    t.Errorf("got %v\nwant %v", actual, 1)
  }

  actual = group.countEmptyCellToLast(5)

  if actual != 1 {
    t.Errorf("got %v\nwant %v", actual, 1)
  }
}
