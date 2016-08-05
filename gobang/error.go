package gobang

func NewCellNotFoundError(point *Point) CellNotFoundError {
  return CellNotFoundError {
    Message: "You have specified a not exist cell",
    Point: point,
  }
}

func NewAlreadyPlacedError(point *Point) AlreadyPlacedError {
  return AlreadyPlacedError {
    Message: "Already the stone is placed",
    Point: point,
  }
}

type CellNotFoundError struct {
  Message string `json:"message"`
  Point *Point `json:"point"`
}

func (c CellNotFoundError) Error() string {
  return c.Message
}

type AlreadyPlacedError struct {
  Message string `json:"message"`
  Point *Point `json:"point"`
}

func (c AlreadyPlacedError) Error() string {
  return c.Message
}
