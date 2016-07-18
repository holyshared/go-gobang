package game

type VerticalSelector struct {
  count int
  board *Board
}

func NewVerticalSelector(board *Board, count int) VerticalSelector {
  return VerticalSelector { count: count, board: board }
}

func (selector *VerticalSelector) At(x, y int) *Cell {
  return selector.board.Select(x, y)
}

func (selector *VerticalSelector) Select(stone Stone) *MatchedResult {
  result := MatchedResult {}

  for x := 0; x <= selector.board.Width() - 1; x++ {
    reachedResult := selector.SelectByVertical(x, stone)

    if reachedResult.IsEmpty() {
      continue
    }
    result.results = append(result.results, reachedResult)
  }

  return &result
}

func (selector *VerticalSelector) SelectByVertical(x int, stone Stone) *ReachedResult {
  result := ReachedResult {}

  for y := 0; y <= selector.board.Height() - 1; y++ {
    cell := selector.At(x, y)

    if cell.Have(stone) == false {
      result.cells = result.cells[0:0]
      continue
    }
    result.cells = append(result.cells, cell)

    if (len(result.cells) < selector.count) {
      continue
    }
    first := result.First()

    if selector.board.Have(x, first.y - 1) {
      p := selector.At(x, first.y - 1)
      result.neighborCells = append(result.neighborCells, p)
    }
    last := result.Last()

    if selector.board.Have(x, last.y + 1) {
      n := selector.At(x, last.y + 1)
      result.neighborCells = append(result.neighborCells, n)
    }
    break
  }

  return &result
}
