package gobang

import (
  "strconv"
  "encoding/json"
)

func NewPoint(x, y int) *Point {
  return &Point {
    x: x,
    y: y,
  }
}

type Point2D interface {
  X() int
  Y() int
  SetTo(x, y int) Point2D
  String() string
}

type Point struct {
  x int `json:"x"`
  y int `json:"y"`
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
