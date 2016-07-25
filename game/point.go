package game

import (
  "strconv"
)

func NewPoint(x, y int) Point {
  return Point {
    x: x,
    y: y,
  }
}

type Point struct {
  x int `json:"x"`
  y int `json:"y"`
}

func (point Point) ToString() string {
  return strconv.Itoa(point.x) + ":" + strconv.Itoa(point.y)
}
