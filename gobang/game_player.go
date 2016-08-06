package gobang

type GamePlayer struct {
  *GobangPlayer
}

func NewGamePlayer(stone Stone) *GamePlayer {
  player := NewGobangPlayer(stone)

  return &GamePlayer {
    GobangPlayer: player,
  }
}
