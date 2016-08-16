package gobang

func NewNpcPlayer(stone Stone, ai GobangAI, selector CellSelector) *NpcPlayer {
  return &NpcPlayer {
    GobangPlayer: NewGobangPlayer(stone, selector),
    GobangAI: ai,
  }
}

type NpcPlayer struct {
  *GobangPlayer
  GobangAI
}
