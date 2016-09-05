package gobang

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
	board := NewBoard(NewSize(30, 30))

	for x := 29; x >= 4; x-- {
		TopRightAxisFillBoard(board, x, 0)
	}
	for x := 23; x >= 4; x-- {
		TopRightAxisFillBoard(board, x, 6)
	}
	TopRightAxisFillBoard(board, 29, 1)
	TopRightAxisFillBoard(board, 23, 7)

	board.Print()

	selector := NewTopRightDiagonalCellMatcher(Black, 5)
	result := selector.Matches(board)

	if len(result.results) != 48 {
		t.Errorf("got %v\nwant %v", len(result.results), 48)
	}
}

func TopRightAxisFillBoard(board *Board, startX int, startY int) {
	y := startY

	for x := startX; x >= 0; x-- {
		if y > startY+4 {
			break
		}
		cell := board.SelectCell(NewPoint(x, y))
		Black.PutTo(cell)
		y++
	}
}

func TestTopRightDiagonalSelectorScanCellGroup(t *testing.T) {
	index := map[string]int{}
	sboard := NewBoard(NewSize(10, 10))
	vboard := NewBoard(NewSize(10, 10))

	selector := NewTopRightDiagonalCellMatcher(Black, 5)
	groups := selector.scanAllCellGroup(sboard)

	if len(groups) != 11 {
		t.Errorf("got %v\nwant %v", len(groups), 11)
	}

	for _, group := range groups {
		for _, v := range group.cells {
			cell := vboard.SelectCell(NewPoint(v.X(), v.Y()))
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
