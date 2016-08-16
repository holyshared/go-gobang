package gobang

type CellSelector interface {
  HaveCell(point *Point) bool
  SelectCell(point *Point) *Cell
}
