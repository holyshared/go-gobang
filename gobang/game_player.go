package gobang

func NewGamePlayer(stone Stone, selector CellSelector) *GamePlayer {
	return &GamePlayer{
		GobangPlayer: NewGobangPlayer(stone),
		selector:     selector,
	}
}

type GamePlayer struct {
	*GobangPlayer
	selector CellSelector
}

func (player *GamePlayer) SelectBoardCell(point Point2D) (*Cell, error) {
	if !player.selector.HaveCell(point) {
		return nil, NewCellNotFoundError(point)
	}
	return player.selector.SelectCell(point), nil
}
