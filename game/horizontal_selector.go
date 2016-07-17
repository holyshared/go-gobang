package game

type HorizontalSelector struct {
  count int
  board *Board
}

func NewHorizontalSelector(board *Board, count int) HorizontalSelector {
  return HorizontalSelector { count: count, board: board }
}

func (selector *HorizontalSelector) At(x, y int) *Cell {
  return selector.board.At(x, y)
}

func (selector *HorizontalSelector) Select(stone Stone) *MatchedResult {
  result := MatchedResult {}

  for y := 0; y <= selector.board.Height() - 1; y++ {
    reachedResult := selector.SelectByHorizontal(y, stone)

    if reachedResult.IsEmpty() {
      continue
    }
    result.results = append(result.results, reachedResult)
  }

  return &result
}

func (selector *HorizontalSelector) SelectByHorizontal(y int, stone Stone) *ReachedResult {
  result := ReachedResult {}

  for x := 0; x <= selector.board.Width() - 1; x++ {
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

    if first.point.x > 0 {
      p := selector.At(first.point.x - 1, y)
      result.neighborCells = append(result.neighborCells, p)
    }
    last := result.Last()

    if last.point.x < selector.board.Width() - 1 {
      n := selector.At(last.point.x + 1, y)
      result.neighborCells = append(result.neighborCells, n)
    }
    break
  }

  return &result
}
