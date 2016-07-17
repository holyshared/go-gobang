package game

import (
  "testing"
)

func TestSelectorTopFilled(t *testing.T) {
  board := NewBoard(10, 10)
  for y := 0; y <= 4; y++ {
    for x := 0; x <= board.Width() - 1; x++ {
      cell := board.At(x, y)
      Black.PutTo(cell)
    }
  }
  board.Print()

  selector := NewVerticalSelector(&board, 5)
  result := selector.Select(Black)

  if len(result.results) != 10 {
    t.Errorf("got %v\nwant %v", len(result.results), 10)
  }
}

func TestSelectorBottomFilled(t *testing.T) {
  board := NewBoard(10, 10)
  for y := 5; y <= board.Height() - 1; y++ {
    for x := 0; x <= board.Width() - 1; x++ {
      cell := board.At(x, y)
      Black.PutTo(cell)
    }
  }
  board.Print()

  selector := NewVerticalSelector(&board, 5)
  result := selector.Select(Black)

  if len(result.results) != 10 {
    t.Errorf("got %v\nwant %v", len(result.results), 10)
  }
}
