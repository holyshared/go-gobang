package game

import (
  "fmt"
  "strings"
)

func NewBoard(height, width int) Board {
  size := Size { height: height, width: width }
  cells := make([]Cell, size.CellCount())

  for i := range cells {
    y := (i / size.width)
    x := i - (y * size.width)
    cells[i] = Cell { Point { x: x, y: y }, 0 }
  }
  return Board { size, cells }
}

type Board struct {
  Size
  cells []Cell
}

func (board *Board) Select(x, y int) *Cell {
  index := (y * board.width) + x
  return &board.cells[index]
}

func (board *Board) Have(x, y int) bool {
  isXRange := x >= 0 && x <= board.width - 1
  isYRange := y >= 0 && y <= board.height - 1
  return isXRange && isYRange
}

func (board *Board) Height() int {
  return board.height
}

func (board *Board) Width() int {
  return board.width
}

func (board *Board) IsCellEmpty(x, y int) bool {
  cell := board.Select(x, y)
  return cell.IsEmpty()
}

func (board *Board) IsAllFilled() bool {
  endX := board.Width() - 1
  endY := board.Height() - 1

  for x := 0; x <= endX; x++ {
    for y := 0; y <= endY; y++ {
      if !board.IsCellEmpty(x, y) {
        continue
      }
      return false
    }
  }

  return true
}

func (board *Board) Print() {
  cells := make([]string, 0)

  for y := 0; y <= board.Height() - 1; y++ {
    for x := 0; x <= board.Width() - 1; x++ {
      cell := board.Select(x, y)

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
