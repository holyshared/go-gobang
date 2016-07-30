package gobang

import (
  "testing"
)

/**
 * |B|B|B|B|B|B| | | | |
 * |B|B|B|B|B|B|B| | | |
 * |B|B|B|B|B|B|B|B| | |
 * |B|B|B|B|B|B|B|B|B| |
 * |B|B|B|B|B|B|B|B|B|B|
 * |B|B|B|B|B|B|B|B|B|B|
 * | |B|B|B|B|B|B|B|B|B|
 * | | |B|B|B|B|B|B|B|B|
 * | | | |B|B|B|B|B|B|B|
 * | | | | |B|B|B|B|B|B|
 */
func TestTopLeftDiagonalSelector(t *testing.T) {
  board := NewBoard(30, 30)

  for x := 0; x <= board.Width() - 5; x++ {
    TopLeftFillBoard(&board, x, 0)
  }
  for x := 6; x <= board.Width() - 5; x++ {
    TopLeftFillBoard(&board, x, 6)
  }

  TopLeftFillBoard(&board, 0, 1)
  TopLeftFillBoard(&board, 6, 7)

  board.Print()

  selector := NewTopLeftDiagonalCellMatcher(Black, 5)
  result := selector.Matches(&board)

  if len(result.results) != 48 {
    t.Errorf("got %v\nwant %v", len(result.results), 48)
  }
}

func TopLeftFillBoard(board *Board, startX int, startY int) {
  y := startY

  for x := startX; x <= board.Width() - 1; x++ {
    if y > startY + 4 {
      break
    }
    cell := board.Select(x, y)
    Black.PutTo(cell)
    y++
  }
}

func TestTopLeftDiagonalSelectorScanCellGroup(t *testing.T) {
  index := map[string]int{}
  sboard := NewBoard(10, 10)
  vboard := NewBoard(10, 10)

  selector := NewTopLeftDiagonalCellMatcher(Black, 5)
  groups := selector.scanAllCellGroup(&sboard)

  if len(groups) != 11 {
    t.Errorf("got %v\nwant %v", len(groups), 11)
  }

  for _, group := range groups {
    for _, v := range group.cells {
      cell := vboard.Select(v.X, v.Y)
      Black.PutTo(cell)
      index[cell.String()] += 1
    }
  }

  for k, v := range index {
    if v <= 1 {
      continue
    }
    t.Errorf("cell is a duplicate %v", k)
  }
  vboard.Print()
}
