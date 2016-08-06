package gobang

func NewSize(height, width int) *Size {
  return &Size {
    Height: height,
    Width: width,
  }
}

type Size struct {
  Height int `json:"height"`
  Width int `json:"width"`
}

func (size *Size) CellCount() int {
  return size.Height * size.Width
}
