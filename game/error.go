package game

func NewCellNotFoundError(point Point) CellNotFoundError {
  return CellNotFoundError {
    message: "You have specified a not exist cell",
    point: point,
  }
}

func NewAlreadyPlacedError(point Point) AlreadyPlacedError {
  return AlreadyPlacedError {
    message: "Already the stone is placed",
    point: point,
  }
}

type CellNotFoundError struct {
  message string `json:"message"`
  point Point `json:"point"`
}

func (c CellNotFoundError) Error() string {
  return c.message
}

type AlreadyPlacedError struct {
  message string `json:"message"`
  point Point `json:"point"`
}

func (c AlreadyPlacedError) Error() string {
  return c.message
}
