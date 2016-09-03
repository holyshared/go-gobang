package gobang

type CellGroup struct {
  cells []*Cell
}

func (g *CellGroup) SelectReached(selector ReachedSelector) []*ReachedResult {
  return selector.Select(g)
}

// XXX Should I use sort package?
func (g *CellGroup) countEmptyCellToFirst(i int) int {
  if i <= 0 {
    return 0
  }

  cells := g.cells[0:i + 1]
  reverseCells := make([]*Cell, len(cells))
  reverseIndex := 0

  for j := len(cells) - 1; j >= 0; j-- {
    reverseCells[reverseIndex] = cells[j]
    reverseIndex++
  }

  return countNeighborEmptyCell(reverseCells)
}

func (g *CellGroup) countEmptyCellToLast(i int) int {
  cells := g.cells[i:]
  return countNeighborEmptyCell(cells)
}

func countNeighborEmptyCell(cells []*Cell) int {
  emptyCount := 0

  for _, cell := range cells {
    if !cell.IsEmpty() {
      break
    }
    emptyCount++
  }

  return emptyCount
}
