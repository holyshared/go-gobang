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

type MatchedResult struct {
  results []*ReachedResult
}

type ReachedResult struct {
  cells []*Cell
  neighborCells []*Cell
}

func (selector *Selector) Select(stone Stone) *MatchedResult {
  cells := make([]*Cell, 0)
  neighborCells := make([]*Cell, 0)

  result := MatchedResult {}

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
      first := cells[0]

      if first.point.x > 0 {
        p := selector.At(first.point.x - 1, y)
        neighborCells = append(neighborCells, p)
      }

      last := cells[len(cells) - 1]

      if last.point.x <= selector.board.Width() - 1 {
        n := selector.At(first.point.x + 1, y)
        neighborCells = append(neighborCells, n)
      }

      reachedResult := ReachedResult { cells: cells[:0], neighborCells: neighborCells[:0] }
      result.results = append(result.results, &reachedResult)

      cells = cells[0:0]

      break
    }
  }

  return &result
}
