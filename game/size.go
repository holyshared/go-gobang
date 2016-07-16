package game

type Size struct {
  height, width uint
}

func (size *Size) CellCount() uint {
  return size.height * size.width
}
