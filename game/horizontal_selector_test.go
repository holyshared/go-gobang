package game

import (
  "testing"
)

func TestSelectorRightFilled(t *testing.T) {
  board := NewBoard(10, 10)
  for y := 0; y <= board.Height() - 1; y++ {
    for x := 5; x <= board.Width() - 1; x++ {
      cell := board.Select(x, y)
      Black.PutTo(cell)
    }
  }
  board.Print()

  selector := NewHorizontalSelector(&board, 5)
  result := selector.Select(Black)

  if len(result.results) != 10 {
    t.Errorf("got %v\nwant %v", len(result.results), 10)
  }
}

func TestSelectorLeftFilled(t *testing.T) {
  board := NewBoard(10, 10)
  for y := 0; y <= board.Height() - 1; y++ {
    for x := 0; x <= 4; x++ {
      cell := board.Select(x, y)
      Black.PutTo(cell)
    }
  }
  board.Print()

  selector := NewHorizontalSelector(&board, 5)
  result := selector.Select(Black)

  if len(result.results) != 10 {
    t.Errorf("got %v\nwant %v", len(result.results), 10)
  }
}
