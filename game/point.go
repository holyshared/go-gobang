package game

import (
  "strconv"
)

type Point struct {
  x, y int
}

func (point Point) ToString() string {
  return strconv.Itoa(point.x) + ":" + strconv.Itoa(point.y)
}
