package game

type Game struct {
  board *Board
  currentPlayer Player
  player Player
  npcPlayer Player
}

func (g *Game) CurrentBoard() *Board {
  return g.board
}

func (g *Game) CurrentPlayer() Player {
  return g.currentPlayer
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
