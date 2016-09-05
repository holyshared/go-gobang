package gobang

const (
	OneSide int = iota + 1
	BothSides
)

func NewReachedResult(cells, emptyNeighborCells []*Cell, continuousEmptyCellCount int) *ReachedResult {
	var reachedType int
	emptyCellCount := len(emptyNeighborCells)

	switch emptyCellCount {
	default:
		reachedType = 0
	case OneSide:
		reachedType = OneSide
	case BothSides:
		reachedType = BothSides
	}

	return &ReachedResult{
		reachedType:              reachedType,
		cells:                    cells,
		emptyNeighborCells:       emptyNeighborCells,
		continuousEmptyCellCount: continuousEmptyCellCount,
	}
}

type ReachedResult struct {
	reachedType              int
	cells                    []*Cell
	emptyNeighborCells       []*Cell
	continuousEmptyCellCount int
}

func (result *ReachedResult) IsEmpty() bool {
	return len(result.cells) <= 0
}

func (result *ReachedResult) IsNeighborEmpty(num int) bool {
	return result.reachedType == num
}

func (result *ReachedResult) HasEmptyNeighborCell() bool {
	return len(result.emptyNeighborCells) > 0
}

func (result *ReachedResult) EmptyNeighborCells() []*Cell {
	return result.emptyNeighborCells
}

func (result *ReachedResult) ContinuousEmptyCellCount() int {
	return result.continuousEmptyCellCount
}
