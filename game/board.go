package game

func NewBoard(height, width uint) Board {
  size := &Size { height: height, width: width }
  cells := make([]Cell, size.CellCount())

  for i := range cells {
    y := (uint(i) / size.width)
    x := uint(i) - (uint(y) * size.width)
    cells[i] = Cell { point: Point { x: x, y: y } }
  }
  return Board { size: size, cells: cells }
}

type Board struct {
  size *Size
  cells []Cell
}

func (board *Board) At(x, y uint) *Cell {
  index := (y * board.size.width) + x
  return &board.cells[index]
}

func (board *Board) Height() uint {
  return board.size.height
}

func (board *Board) Width() uint {
  return board.size.width
}
