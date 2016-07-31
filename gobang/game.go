package gobang

type Game struct {
  board *Board
  currentPlayer Player
  player *GamePlayer
  npcPlayer Player
}

func (g *Game) CurrentBoard() *Board {
  return g.board
}

func (g *Game) CurrentPlayer() Player {
  return g.currentPlayer
}

func (g *Game) GamePlayer() *GamePlayer {
  return g.player
}

func (g *Game) NpcPlayer() Player {
  return g.npcPlayer
}

func (g *Game) ChangeToNextPlayer() {
  var player Player

  if (g.currentPlayer == g.npcPlayer) {
    player = g.player
  } else {
    player = g.npcPlayer
  }

  g.currentPlayer = player
}
