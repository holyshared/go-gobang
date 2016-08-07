package gobang

type CellMatcher interface {
  Matches(cell *Cell) bool
}

func EmptyCell() *EmptyCellMatcher {
  return &EmptyCellMatcher {}
}

type EmptyCellMatcher struct {
}

func (matcher *EmptyCellMatcher) Matches(cell *Cell) bool {
  return cell.IsEmpty()
}
