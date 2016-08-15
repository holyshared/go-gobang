package gobang

type GamePlayer struct {
  *GobangPlayer
}

func NewGamePlayer(stone Stone, game *GameContext) *GamePlayer {
  player := NewGobangPlayer(stone, game)

  return &GamePlayer {
    GobangPlayer: player,
  }
}
