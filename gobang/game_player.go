package gobang

type GamePlayer struct {
  *GobangPlayer
}

func NewGamePlayer(stone Stone, selector CellSelector) *GamePlayer {
  return &GamePlayer {
    GobangPlayer: NewGobangPlayer(stone, selector),
  }
}
