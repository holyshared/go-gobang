package main

import "fmt"

type Point struct {
  x, y uint
}

type Cell struct {
  stone Stone
  point Point
}

func (cell *Cell) IsEmpty() bool {
  return cell.stone == nil
}

type BlackStone struct {
}

func (stone *BlackStone) PutTo(cell *Cell) {
  cell.stone = stone
}

type WhiteStone struct {
}

func (stone *WhiteStone) PutTo(cell *Cell) {
  cell.stone = stone
}

type Stone interface {
  PutTo(*Cell)
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

func NewBoard(height, width uint) Board {
  size := &Size { height: height, width: width }
  cells := make([]Cell, size.CellCount())

  for i := 0; i <= len(cells) - 1; i++ {
    y := (uint(i) / size.width)
    x := uint(i) - (uint(y) * size.width)
    cells[i] = Cell { point: Point { x: x, y: y } }
  }
  return Board { size: size, cells: cells }
}


func main() {
  board := NewBoard(10, 10)
  cell := board.At(0, 0)

  blackStone := BlackStone {}
  blackStone.PutTo(cell)

  whiteStone := WhiteStone {}
  whiteStone.PutTo(cell)

  fmt.Println(board)
}
