package gobang

func NewGobang(playerStone, npcPlayerStone Stone) *Gobang {
  rule := DefaultGameRule()
  board := NewBoard(rule.BoardSize())


  context := NewGameContext(rule)
  context.player = NewGamePlayer(Black, context)


  ctx := &NpcAIContext {
    rule: rule,
    board: board,
    playerStone: playerStone,
    npcPlayerStone: npcPlayerStone,
  }
  ai := NewNpcAI(ctx)

  context.npcPlayer = NewNpcPlayer(White, ai)
  context.currentPlayer = context.player

  return &Gobang {
    GameContext: context,
  }
}

type Gobang struct {
  *GameContext
}
