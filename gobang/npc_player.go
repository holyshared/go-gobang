package gobang

func NewNpcPlayer(stone Stone, game *GameContext) *NpcPlayer {
  ai := NewNpcArtificialIntelligence(game)
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
