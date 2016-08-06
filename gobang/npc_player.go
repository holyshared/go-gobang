package gobang

func NewNpcPlayer(stone Stone) *NpcPlayer {
  ai := NewNpcArtificialIntelligence()
  player := NewGobangPlayer(stone)

  return &NpcPlayer {
    GobangPlayer: player,
    GobangArtificialIntelligence: ai,
  }
}

type NpcPlayer struct {
  *GobangPlayer
  GobangArtificialIntelligence
}
