package gobang

func NewReachedResult(cells, neighborCell, emptyNeighborCells []*Cell) *ReachedResult {
  return &ReachedResult {
    cells: cells,
    neighborCells: neighborCell,
    emptyNeighborCells: emptyNeighborCells,
  }
}

type ReachedResult struct {
  cells []*Cell
  neighborCells []*Cell
  emptyNeighborCells []*Cell
}

func (result *ReachedResult) IsEmpty() bool {
  return len(result.cells) <= 0
}

func (result *ReachedResult) HasEmptyNeighborCell() bool {
  return len(result.emptyNeighborCells) > 0
}

func (result *ReachedResult) EmptyNeighborCells() []*Cell {
  return result.emptyNeighborCells
}
