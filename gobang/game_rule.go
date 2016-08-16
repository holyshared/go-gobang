package gobang

import (
  "encoding/json"
)

func NewGameRule(size *Size, reachedCount int) *GameRule {
  return &GameRule {
    Size: size,
    reachedStoneCount: reachedCount,
  }
}

func DefaultGameRule() *GameRule {
  return NewGameRule(NewSize(30, 30), 5)
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

func (rule *GameRule) MarshalJSON() ([]byte, error) {
  jsonObject := struct {
    Size *Size `json:"size"`
    ReachedStoneCount int `json:"reachedStoneCount"`
  }{
    Size: rule.BoardSize(),
    ReachedStoneCount: rule.ReachedStoneCount(),
  }
  return json.Marshal(jsonObject)
}
