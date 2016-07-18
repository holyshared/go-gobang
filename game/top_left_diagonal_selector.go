package game

/**
 * |B| | | | | | | | | |
 * | |B| | | | | | | | |
 * | | |B| | | | | | | |
 * | | | |B| | | | | | |
 * | | | | |B| | | | | |
 * | | | |B| | | | | | |
 * | | | | |B| | | | | |
 * | | | | | |B| | | | |
 * | | | | | | |B| | | |
 * | | | | | | | |B| | |
 */

type TopLeftDiagonalSelector struct {
  count int
  board *Board
}

func NewTopLeftDiagonalSelector(board *Board, count int) TopLeftDiagonalSelector {
  return TopLeftDiagonalSelector { count: count, board: board }
}

func (s *TopLeftDiagonalSelector) Select(stone Stone) *MatchedResult {
  result := MatchedResult {}

  /**
   * |B|B|B|B|B|B|B|B|B|B|
   * | |B|B|B|B|B|B|B|B|B|
   * | | |B|B|B|B|B|B|B|B|
   * | | | |B|B|B|B|B|B|B|
   * | | | | |B|B|B|B|B|B|
   * | | | | | |B|B|B|B|B|
   * | | | | | | |B|B|B|B|
   * | | | | | | | |B|B|B|
   * | | | | | | | | |B|B|
   * | | | | | | | | | |B|
   */
  for x := 0; x <= s.board.Width() - 1; x++ {
    xAsixResult := s.Scan(x, 0, stone)

    if xAsixResult.IsEmpty() {
      continue
    }
    result.results = append(result.results, xAsixResult)
  }

  /**
   * | | | | | | | | | | |
   * |B| | | | | | | | | |
   * |B|B| | | | | | | | |
   * |B|B|B| | | | | | | |
   * |B|B|B|B| | | | | | |
   * |B|B|B|B|B| | | | | |
   * |B|B|B|B|B|B| | | | |
   * |B|B|B|B|B|B|B| | | |
   * |B|B|B|B|B|B|B|B| | |
   * |B|B|B|B|B|B|B|B|B| |
   */
  for y := 1; y <= s.board.Height() - 1; y++ {
    yAsixResult := s.Scan(0, y, stone)

    if yAsixResult.IsEmpty() {
      continue
    }
    result.results = append(result.results, yAsixResult)
  }

  return &result
}

func (s *TopLeftDiagonalSelector) Scan(startX, startY int, stone Stone) *ReachedResult {
  y := startY
  result := ReachedResult {}

  for x := startX; x <= s.board.Width() - 1; x++ {
    if y == s.board.Height() - 1 {
      break
    }

    cell := s.board.Select(x, y)
    y++

    if cell.Have(stone) == false {
      result.cells = result.cells[0:0]
      continue
    }
    result.cells = append(result.cells, cell)

    if len(result.cells) < s.count {
      continue
    }

    first := result.First()

    if s.board.Have(first.x - 1, first.y - 1) {
      p := s.board.Select(first.x - 1, first.y - 1)
      result.neighborCells = append(result.neighborCells, p)
    }

    last := result.Last()

    if s.board.Have(last.x + 1, last.y + 1) {
      p := s.board.Select(last.x + 1, last.y + 1)
      result.neighborCells = append(result.neighborCells, p)
    }
    break
  }

  return &result
}
