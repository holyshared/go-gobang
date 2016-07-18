package game

import (
  "fmt"
  "strings"
)

func NewBoard(height, width int) Board {
  size := &Size { height: height, width: width }
  cells := make([]Cell, size.CellCount())

  for i := range cells {
    y := (i / size.width)
    x := i - (y * size.width)
    cells[i] = Cell { Point { x: x, y: y }, 0 }
  }
  return Board { size: size, cells: cells }
}

type Board struct {
  size *Size
  cells []Cell
}

func (board *Board) At(x, y int) *Cell {
  index := (y * board.size.width) + x
  return &board.cells[index]
}

func (board *Board) Height() int {
  return board.size.height
}

func (board *Board) Width() int {
  return board.size.width
}

func (board *Board) Print() {
  cells := make([]string, 0)

  for y := 0; y <= board.Height() - 1; y++ {
    for x := 0; x <= board.Width() - 1; x++ {
      cell := board.At(x, y)

      if (cell.IsEmpty()) {
        cells = append(cells, " ")
      } else {
        cells = append(cells, cell.stone.ToString())
      }
    }
    fmt.Println("|" + strings.Join(cells, "|") + "|")
    cells = cells[0:0]
  }
}
