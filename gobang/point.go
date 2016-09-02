package gobang

import (
  "strconv"
)

func NewPoint(x, y int) *Point {
  return &Point {
    X: x,
    Y: y,
  }
}

type Point2D interface {
  X() int
  Y() int
  SetTo(x, y int) Point2D
  String() string
}

type Point struct {
  X int `json:"x"`
  Y int `json:"y"`
}

func (point *Point) SetTo(x, y int) *Point {
  point.X = x
  point.Y = y
  return point
}

func (point *Point) String() string {
  return strconv.Itoa(point.X) + ":" + strconv.Itoa(point.Y)
}
