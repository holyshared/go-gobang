package game

func NewTopRightNeighborDistance() NeighborDistance {
  prev := PointDistance { x: 1,  y: -1 }
  next := PointDistance { x: -1, y: 1 }

  return NeighborDistance {
    prev: prev,
    next: next,
  }
}

func NewTopLeftNeighborDistance() NeighborDistance {
  prev := PointDistance { x: -1,  y: -1 }
  next := PointDistance { x: 1, y: 1 }

  return NeighborDistance {
    prev: prev,
    next: next,
  }
}

func NewHorizontalNeighborDistance() NeighborDistance {
  prev := PointDistance { x: -1,  y: 0 }
  next := PointDistance { x: 1, y: 0 }

  return NeighborDistance {
    prev: prev,
    next: next,
  }
}

func NewVerticalNeighborDistance() NeighborDistance {
  prev := PointDistance { x: 0,  y: -1 }
  next := PointDistance { x: 0, y: 1 }

  return NeighborDistance {
    prev: prev,
    next: next,
  }
}

type PointDistance struct {
  x, y int
}

type NeighborDistance struct {
  prev PointDistance
  next PointDistance
}

func (n NeighborDistance) prevPoint(cell *Cell) Point {
  return Point {
    x: cell.x + n.prev.x,
    y: cell.y + n.prev.y,
  }
}

func (n NeighborDistance) nextPoint(cell *Cell) Point {
  return Point {
    x: cell.x + n.next.x,
    y: cell.y + n.next.y,
  }
}
