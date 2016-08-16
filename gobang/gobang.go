package gobang

func NewGobang() *Gobang {
  context := NewGameContext(DefaultGameRule())
  context.player = NewGamePlayer(Black, context)

  ai := NewNpcAI(context)
  context.npcPlayer = NewNpcPlayer(White, ai)
  context.currentPlayer = context.player

  return &Gobang {
    GameContext: context,
  }
}

type Gobang struct {
  *GameContext
}
