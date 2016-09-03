package gobang

type ReachedResultBuilder struct {
  count int
  cells []*Cell
  neighborCells []*Cell
}

func (result *ReachedResultBuilder) FirstCell() *Cell {
  if len(result.cells) <= 0 {
    return nil
  }
  return result.cells[0]
}

func (result *ReachedResultBuilder) LastCell() *Cell {
  if len(result.cells) <= 0 {
    return nil
  }
  return result.cells[len(result.cells) - 1]
}

func (result *ReachedResultBuilder) IsReached() bool {
  return len(result.cells) >= result.count
}

func (result *ReachedResultBuilder) AddCell(cell *Cell) {
  result.cells = append(result.cells, cell)
}

func (result *ReachedResultBuilder) AddNeighborCell(cell *Cell) {
  result.neighborCells = append(result.neighborCells, cell)
}

func (result *ReachedResultBuilder) Clear() {
  result.cells = result.cells[0:0]
  result.neighborCells = result.neighborCells[0:0]
}

func (result *ReachedResultBuilder) emptyNeighborCells() []*Cell {
  cells := make([]*Cell, 0)

  for _, cell := range result.neighborCells {
    if !cell.IsEmpty() {
      continue
    }
    cells = append(cells, cell)
  }
  return cells
}

func (result *ReachedResultBuilder) ReachedResult() *ReachedResult {
  return NewReachedResult(
    result.cells,
    result.emptyNeighborCells(),
  )
}
