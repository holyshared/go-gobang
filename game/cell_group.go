package game

type CellGroup struct {
  cells []*Cell
}

func (g *CellGroup) SelectReached(selector ReachedSelector) []*ReachedResult {
  return selector.Select(g)
}
