package gobang

func NewSize(height, width int) *Size {
  return &Size {
    height: height,
    width: width,
  }
}

type Size struct {
  height, width int
}

func (size *Size) Height() int {
  return size.height
}

func (size *Size) Width() int {
  return size.width
}

func (size *Size) CellCount() int {
  return size.height * size.width
}
