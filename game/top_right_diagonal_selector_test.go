package game

import (
  "testing"
)

/**
 * | | | | |B|B|B|B|B|B|
 * | | | |B|B|B|B|B|B|B|
 * | | |B|B|B|B|B|B|B|B|
 * | |B|B|B|B|B|B|B|B|B|
 * |B|B|B|B|B|B|B|B|B|B|
 * |B|B|B|B|B|B|B|B|B|B|
 * |B|B|B|B|B|B|B|B|B| |
 * |B|B|B|B|B|B|B|B| | |
 * |B|B|B|B|B|B|B| | | |
 * |B|B|B|B|B|B| | | | |
 */
func TestTopRightDiagonalSelector(t *testing.T) {
  board := NewBoard(10, 10)

  for x := board.Width() - 1; x >= 4; x-- {
    TopRightXAxisFillBoard(&board, x)
  }
  for y := 1; y <= 5; y++ {
    TopRightYAxisFillBoard(&board, y)
  }
  board.Print()

  selector := NewTopRightDiagonalSelector(&board, 5)
  result := selector.Select(Black)

  if len(result.results) != 11 {
    t.Errorf("got %v\nwant %v", len(result.results), 11)
  }
}

func TopRightXAxisFillBoard(board *Board, startX int) {
  y := 0

  for x := startX; x >= 0; x-- {
    if y > board.Height() - 1 {
      break
    }
    cell := board.Select(x, y)
    Black.PutTo(cell)
    y++
  }
}

func TopRightYAxisFillBoard(board *Board, startY int) {
  x := board.Width() - 1

  for y := startY; y <= board.Height() - 1; y++ {
    if x <= 0 {
      break
    }
    cell := board.Select(x, y)
    Black.PutTo(cell)
    x--
  }
}
