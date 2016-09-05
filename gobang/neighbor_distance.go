package gobang

func NewTopRightNeighborDistance() NeighborDistance {
	prev := PointDistance{x: 1, y: -1}
	next := PointDistance{x: -1, y: 1}

	return NeighborDistance{
		prev: prev,
		next: next,
	}
}

func NewTopLeftNeighborDistance() NeighborDistance {
	prev := PointDistance{x: -1, y: -1}
	next := PointDistance{x: 1, y: 1}

	return NeighborDistance{
		prev: prev,
		next: next,
	}
}

func NewHorizontalNeighborDistance() NeighborDistance {
	prev := PointDistance{x: -1, y: 0}
	next := PointDistance{x: 1, y: 0}

	return NeighborDistance{
		prev: prev,
		next: next,
	}
}

func NewVerticalNeighborDistance() NeighborDistance {
	prev := PointDistance{x: 0, y: -1}
	next := PointDistance{x: 0, y: 1}

	return NeighborDistance{
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

func (n NeighborDistance) prevPoint(cell *Cell) Point2D {
	return NewPoint(
		cell.X()+n.prev.x,
		cell.Y()+n.prev.y,
	)
}

func (n NeighborDistance) nextPoint(cell *Cell) Point2D {
	return NewPoint(
		cell.X()+n.next.x,
		cell.Y()+n.next.y,
	)
}
