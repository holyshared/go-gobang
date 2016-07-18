package game

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
  board := NewBoard(10, 10)

  for y := 0; y <= 5; y++ {
    TopLeftYAxisFillBoard(&board, y)
  }
  for x := 1; x <= 5; x++ {
    TopLeftXAxisFillBoard(&board, x)
  }
  board.Print()

  selector := NewTopLeftDiagonalSelector(&board, 5)
  result := selector.Select(Black)

  if len(result.results) != 11 {
    t.Errorf("got %v\nwant %v", len(result.results), 11)
  }
}

func TopLeftYAxisFillBoard(board *Board, startY int) {
  y := startY

  for x := 0; x <= board.Width() - startY; x++ {
    if y > board.Height() - 1 {
      break
    }
    cell := board.Select(x, y)
    Black.PutTo(cell)
    y++
  }
}

func TopLeftXAxisFillBoard(board *Board, startX int) {
  x := startX

  for y := 0; y <= board.Height() - startX; y++ {
    if x > board.Width() - 1 {
      break
    }
    cell := board.Select(x, y)
    Black.PutTo(cell)
    x++
  }
}
