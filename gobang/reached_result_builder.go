package gobang

type ReachedResultBuilder struct {
	count                    int
	cells                    []*Cell
	firstNeighborCell        *Cell
	lastNeighborCell         *Cell
	continuousEmptyCellCount int
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
	return result.cells[len(result.cells)-1]
}

func (result *ReachedResultBuilder) IsReached() bool {
	return len(result.cells) >= result.count
}

func (result *ReachedResultBuilder) AddCell(cell *Cell) {
	result.cells = append(result.cells, cell)
}

func (result *ReachedResultBuilder) Clear() {
	result.cells = result.cells[0:0]
	result.firstNeighborCell = nil
	result.lastNeighborCell = nil
}

func (result *ReachedResultBuilder) SetFirstNeighborCell(cell *Cell) {
	result.firstNeighborCell = cell
}

func (result *ReachedResultBuilder) SetLastNeighborCell(cell *Cell) {
	result.lastNeighborCell = cell
}

func (result *ReachedResultBuilder) SetContinuousEmptyCellCount(count int) {
	result.continuousEmptyCellCount = count
}

func (result *ReachedResultBuilder) emptyNeighborCells() []*Cell {
	cells := make([]*Cell, 0)
	neighborCells := []*Cell{
		result.firstNeighborCell,
		result.lastNeighborCell,
	}

	for _, cell := range neighborCells {
		if cell == nil {
			continue
		}
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
		result.continuousEmptyCellCount,
	)
}
