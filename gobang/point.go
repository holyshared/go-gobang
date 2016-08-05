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

type Point struct {
  X int `json:"x"`
  Y int `json:"y"`
}

func (point *Point) String() string {
  return strconv.Itoa(point.X) + ":" + strconv.Itoa(point.Y)
}
