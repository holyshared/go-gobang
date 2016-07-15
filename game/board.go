package game

type Point struct {
  x, y uint
}

type Cell struct {
  stone Stone
  point Point
}

func (cell *Cell) IsEmpty() bool {
  return cell.stone == None
}

type Size struct {
  height, width uint
}

func (size *Size) CellCount() uint {
  return size.height * size.width
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

func NewBoard(height, width uint) Board {
  size := &Size { height: height, width: width }
  cells := make([]Cell, size.CellCount())

  for i := range cells {
    y := (uint(i) / size.width)
    x := uint(i) - (uint(y) * size.width)
    cells[i] = Cell { stone: None, point: Point { x: x, y: y } }
  }
  return Board { size: size, cells: cells }
}
