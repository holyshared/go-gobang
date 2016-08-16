package gobang

func NewGobang(rule *GameRule, playerStone, npcPlayerStone Stone) *Gobang {
  context := NewGameContext(rule, playerStone, npcPlayerStone)

  return &Gobang {
    GameContext: context,
  }
}

type Gobang struct {
  *GameContext
}
