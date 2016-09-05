package gobang

import (
	"testing"
)

func TestSelectorTopFilled(t *testing.T) {
	board := NewBoard(NewSize(30, 30))

	for y := 0; y <= 4; y++ {
		for x := 0; x <= board.Width()-1; x++ {
			cell := board.SelectCell(NewPoint(x, y))
			Black.PutTo(cell)
		}
	}

	for y := 6; y <= 10; y++ {
		for x := 0; x <= board.Width()-1; x++ {
			cell := board.SelectCell(NewPoint(x, y))
			Black.PutTo(cell)
		}
	}

	board.Print()

	selector := NewVerticalCellMatcher(Black, 5)
	result := selector.Matches(board)

	if len(result.results) != 60 {
		t.Errorf("got %v\nwant %v", len(result.results), 60)
	}
}

func TestSelectorBottomFilled(t *testing.T) {
	board := NewBoard(NewSize(30, 30))

	for y := 19; y <= 23; y++ {
		for x := 0; x <= board.Width()-1; x++ {
			cell := board.SelectCell(NewPoint(x, y))
			Black.PutTo(cell)
		}
	}

	for y := 25; y <= board.Height()-1; y++ {
		for x := 0; x <= board.Width()-1; x++ {
			cell := board.SelectCell(NewPoint(x, y))
			Black.PutTo(cell)
		}
	}
	board.Print()

	selector := NewVerticalCellMatcher(Black, 5)
	result := selector.Matches(board)

	if len(result.results) != 60 {
		t.Errorf("got %v\nwant %v", len(result.results), 60)
	}
}
