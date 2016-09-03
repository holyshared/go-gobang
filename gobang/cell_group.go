package gobang

type CellGroup struct {
  cells []*Cell
}

func (g *CellGroup) SelectReached(selector ReachedSelector) []*ReachedResult {
  return selector.Select(g)
}

func (g *CellGroup) countEmptyCellToFirst(i int) int {
  cells := g.cells[0:i + 1]
  return countNeighborEmptyCell(cells)
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
