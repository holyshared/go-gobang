package gobang

type CellSelector interface {
  HaveCell(point Point2D) bool
  SelectCell(point Point2D) *Cell
}
