package gobang

func NewNpcPlayer(stone Stone, ai GobangArtificialIntelligence, selector CellSelector) *NpcPlayer {
  return &NpcPlayer {
    GobangPlayer: NewGobangPlayer(stone, selector),
    GobangArtificialIntelligence: ai,
  }
}

type NpcPlayer struct {
  *GobangPlayer
  GobangArtificialIntelligence
}
