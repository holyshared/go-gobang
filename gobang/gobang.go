package gobang

func NewGobang() *Gobang {
  context := NewGameContext(DefaultGameRule())
  context.player = NewGamePlayer(Black, context)

  ai := NewNpcArtificialIntelligence(context)
  context.npcPlayer = NewNpcPlayer(White, ai, context)
  context.currentPlayer = context.player

  return &Gobang {
    GameContext: context,
  }
}

type Gobang struct {
  *GameContext
}
