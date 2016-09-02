package gobang

import (
  "strconv"
  "encoding/json"
)

func NewPoint(x, y int) Point2D {
  return &Point {
    x: x,
    y: y,
  }
}

func DefaultPoint() Point2D {
  return NewPoint(0, 0)
}

type Point2D interface {
  X() int
  Y() int
  SetTo(x, y int) Point2D
  String() string
}

type Point struct {
  x int
  y int
}

func (point *Point) X() int {
  return point.x
}

func (point *Point) Y() int {
  return point.y
}

func (point *Point) SetTo(x, y int) Point2D {
  point.x = x
  point.y = y
  return point
}

func (point *Point) String() string {
  return strconv.Itoa(point.x) + ":" + strconv.Itoa(point.y)
}

func (point *Point) MarshalJSON() ([]byte, error) {
  object := struct {
    X int `json:"x"`
    Y int `json:"y"`
  }{
    X: point.X(),
    Y: point.Y(),
  }
  return json.Marshal(object)
}

func (point *Point) UnmarshalJSON(data []byte) error {
  object := struct {
    X int `json:"x"`
    Y int `json:"y"`
  }{}
  err := json.Unmarshal(data, &object)

  if err != nil {
    return err
  }
  point.SetTo(object.X, object.Y)
  return nil
}
