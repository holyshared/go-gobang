package gobang

type NpcPlayer struct {
  *GobangPlayer
}

func (npc *NpcPlayer) SelectTargetCell() *Cell {
  ai := NpcArtificialIntelligence {
    game: npc.game,
  }
  return ai.SelectTargetCell()
}
