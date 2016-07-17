package game

type Selector struct {
  count int
  board *Board
}

func NewSelector(board *Board, count int) Selector {
  return Selector { count: count, board: board }
}

func (selector *Selector) At(x, y int) *Cell {
  return selector.board.At(x, y)
}

func (selector *Selector) Select(stone Stone) [][]*Cell {
  cells := make([]*Cell, 0)
  results := make([][]*Cell, 0)

  for y := 0; y <= selector.board.Height() - 1; y++ {
    for x := 0; x <= selector.board.Width() - 1; x++ {
      cell := selector.At(x, y)

      if cell.Have(stone) == false {
        cells = cells[0:0]
        continue
      }
      cells = append(cells, cell)

      if (len(cells) < selector.count) {
        continue
      }
      results = append(results, cells[:0])
      cells = cells[0:0]

      break
    }
  }

  return results
}
