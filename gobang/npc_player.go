package gobang

func NewNpcPlayer(stone Stone, ai GobangAI) *NpcPlayer {
	return &NpcPlayer{
		GobangPlayer: NewGobangPlayer(stone),
		GobangAI:     ai,
	}
}

type NpcPlayer struct {
	*GobangPlayer
	GobangAI
}
