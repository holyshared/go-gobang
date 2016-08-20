package gobang

import (
  "encoding/json"
)

func NewGobang(rule *GameRule, playerStone, npcPlayerStone Stone) *Gobang {
  context := NewGameContext(rule, playerStone, npcPlayerStone)
  facilitator := NewGameFacilitator(context)

  return &Gobang {
//    GameContext: context,
    GameFacilitator: facilitator,
  }
}

type Gobang struct {
  //*GameContext
  *GameFacilitator
}

func (g *Gobang) MarshalJSON() ([]byte, error) {
  return json.Marshal(g.GameFacilitator.ctx)
}
