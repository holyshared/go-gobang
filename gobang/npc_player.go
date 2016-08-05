package gobang

func NewNpcPlayer(stone Stone) *NpcPlayer {
  ai := NewNpcArtificialIntelligence()
  player := NewGobangPlayer(stone)

  return &NpcPlayer {
    player,
    ai,
  }
}

type NpcPlayer struct {
  *GobangPlayer
  GobangArtificialIntelligence
}
