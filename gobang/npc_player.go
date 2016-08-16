package gobang

func NewNpcPlayer(stone Stone, ai GobangArtificialIntelligence, game *GameContext) *NpcPlayer {
  player := NewGobangPlayer(stone, game)

  return &NpcPlayer {
    GobangPlayer: player,
    GobangArtificialIntelligence: ai,
  }
}

type NpcPlayer struct {
  *GobangPlayer
  GobangArtificialIntelligence
}
