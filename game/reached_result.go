package game

type ReachedResult struct {
  cells []*Cell
  neighborCells []*Cell
}

func (result *ReachedResult) IsEmpty() bool {
  return len(result.cells) <= 0
}

func (result *ReachedResult) First() *Cell {
  if len(result.cells) <= 0 {
    return nil
  }
  return result.cells[0]
}

func (result *ReachedResult) Last() *Cell {
  if len(result.cells) <= 0 {
    return nil
  }
  return result.cells[len(result.cells) - 1]
}

func (result *ReachedResult) IsReached(count int) bool {
  return len(result.cells) >= count
}

func (result *ReachedResult) AddCell(cell *Cell) {
  result.cells = append(result.cells, cell)
}

func (result *ReachedResult) AddNeighborCell(cell *Cell) {
  result.neighborCells = append(result.neighborCells, cell)
}

func (result *ReachedResult) Clear() {
  result.cells = result.cells[0:0]
  result.neighborCells = result.neighborCells[0:0]
}

func (result *ReachedResult) EmptyNeighborCells() []*Cell {
  cells := make([]*Cell, 0)

  for _, cell := range result.neighborCells {
    if !cell.IsEmpty() {
      continue
    }
    cells = append(cells, cell)
  }
  return cells
}
