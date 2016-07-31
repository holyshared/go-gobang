package gobang

type GameContext struct {
  board *Board
  currentPlayer Player
  player *GamePlayer
  npcPlayer *NpcPlayer
}

func (g *GameContext) CurrentBoard() *Board {
  return g.board
}

func (g *GameContext) CurrentPlayer() Player {
  return g.currentPlayer
}

func (g *GameContext) GamePlayer() *GamePlayer {
  return g.player
}

func (g *GameContext) NpcPlayer() *NpcPlayer {
  return g.npcPlayer
}

func (g *GameContext) ChangeToNextPlayer() {
  var player Player

  if (g.currentPlayer == g.npcPlayer) {
    player = g.player
  } else {
    player = g.npcPlayer
  }

  g.currentPlayer = player
}
