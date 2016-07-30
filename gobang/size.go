package gobang

type Size struct {
  height, width int
}

func (size *Size) CellCount() int {
  return size.height * size.width
}
