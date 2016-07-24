package game

import (
  "testing"
)

func TestSelectorRightFilled(t *testing.T) {
  board := NewBoard(30, 30)
  for y := 0; y <= board.Height() - 1; y++ {
    for x := 25; x <= board.Width() - 1; x++ {
      cell := board.Select(x, y)
      Black.PutTo(cell)
    }
    for x := 19; x <= 23; x++ {
      cell := board.Select(x, y)
      Black.PutTo(cell)
    }
  }
  board.Print()

  selector := NewHorizontalCellMatcher(Black, 5)
  result := selector.Matches(&board)

  if len(result.results) != 60 {
    t.Errorf("got %v\nwant %v", len(result.results), 60)
  }
}

func TestSelectorLeftFilled(t *testing.T) {
  board := NewBoard(30, 30)
  for y := 0; y <= board.Height() - 1; y++ {
    for x := 0; x <= 4; x++ {
      cell := board.Select(x, y)
      Black.PutTo(cell)
    }
    for x := 6; x <= 11; x++ {
      cell := board.Select(x, y)
      Black.PutTo(cell)
    }
  }
  board.Print()

  selector := NewHorizontalCellMatcher(Black, 5)
  result := selector.Matches(&board)

  if len(result.results) != 60 {
    t.Errorf("got %v\nwant %v", len(result.results), 60)
  }
}
