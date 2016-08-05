package gobang

import (
  "fmt"
  "strings"
)

func NewBoard(size *Size) *Board {
  cells := make([]Cell, size.CellCount())

  for i := range cells {
    y := (i / size.width)
    x := i - (y * size.width)
    cells[i] = Cell { Point { X: x, Y: y }, 0 }
  }
  return &Board { size, cells }
}

type Board struct {
  *Size
  cells []Cell
}

func (board *Board) Select(point Point) *Cell {
  index := (point.Y * board.width) + point.X
  return &board.cells[index]
}

func (board *Board) Have(point Point) bool {
  isXRange := point.X >= 0 && point.X <= board.width - 1
  isYRange := point.Y >= 0 && point.Y <= board.height - 1
  return isXRange && isYRange
}

func (board *Board) Height() int {
  return board.height
}

func (board *Board) Width() int {
  return board.width
}

func (board *Board) IsCellEmpty(point Point) bool {
  cell := board.Select(point)
  return cell.IsEmpty()
}

func (board *Board) IsAllFilled() bool {
  var point Point
  endX := board.Width() - 1
  endY := board.Height() - 1

  for x := 0; x <= endX; x++ {
    for y := 0; y <= endY; y++ {
      point.X = x
      point.Y = y

      if !board.IsCellEmpty(point) {
        continue
      }
      return false
    }
  }

  return true
}

func (board *Board) Print() {
  var point Point
  cells := make([]string, 0)

  for y := 0; y <= board.Height() - 1; y++ {
    for x := 0; x <= board.Width() - 1; x++ {
      point.X = x
      point.Y = y

      cell := board.Select(point)

      if (cell.IsEmpty()) {
        cells = append(cells, " ")
      } else {
        cells = append(cells, cell.stone.String())
      }
    }
    fmt.Println("|" + strings.Join(cells, "|") + "|")
    cells = cells[0:0]
  }
}
