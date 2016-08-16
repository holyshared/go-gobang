package gobang

func NewGobang(rule *GameRule, playerStone, npcPlayerStone Stone) *Gobang {
  context := NewGameContext(rule, playerStone, npcPlayerStone)
  facilitator := NewGameFacilitator(context)

  return &Gobang {
    GameFacilitator: facilitator,
  }
}

type Gobang struct {
  *GameFacilitator
}
