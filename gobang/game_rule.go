package gobang

func NewGameRule(size *Size, reachedCount int) *GameRule {
  return &GameRule {
    Size: size,
    reachedStoneCount: reachedCount,
  }
}

type GameRule struct {
  *Size
  reachedStoneCount int
}

func (rule *GameRule) BoardSize() *Size {
  return rule.Size
}

func (rule *GameRule) ReachedStoneCount() int {
  return rule.reachedStoneCount
}
