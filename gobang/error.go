package gobang

func NewCellNotFoundError(point Point2D) CellNotFoundError {
	return CellNotFoundError{
		Message: "You have specified a not exist cell",
		Point:   point,
	}
}

func NewAlreadyPlacedError(cell *Cell) AlreadyPlacedError {
	return AlreadyPlacedError{
		Message: "Already the stone is placed",
		Point:   cell.Point(),
	}
}

type CellNotFoundError struct {
	Message string  `json:"message"`
	Point   Point2D `json:"point"`
}

func (c CellNotFoundError) Error() string {
	return c.Message
}

type AlreadyPlacedError struct {
	Message string  `json:"message"`
	Point   Point2D `json:"point"`
}

func (c AlreadyPlacedError) Error() string {
	return c.Message
}
