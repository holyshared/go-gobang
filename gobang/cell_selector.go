package gobang

type CellSelector interface {
  SelectCell(point *Point) (*Cell, error)
}
