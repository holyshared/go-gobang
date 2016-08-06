package gobang

type GameFacilitator struct {
  game *GameContext
}

func (f *GameFacilitator) PlayerSelectCell(point *Point) (*Cell, error) {
  player := f.game.GamePlayer()
  cell, err := player.SelectBoardCell(point)

  return cell, err
}

func (f *GameFacilitator) PlayerPutStoneTo(cell *Cell) (GameResult, error) {
  player := f.game.GamePlayer()
  err := player.PutStoneTo(cell)

  if err != nil {
    return 0, err // FIXME GameResult
  }
  result := f.game.CheckBoard()

  if result == Reached {
    return Win, nil
  } else if result == Filled {
    return Draw, nil
  } else {
    return 0, nil // FIXME GameResult
  }
}

func (f *GameFacilitator) ChangeToNextPlayer() {
  f.game.ChangeToNextPlayer()
}

func (f *GameFacilitator) NpcPlayerPutStone() (GameResult, error) {
  player := f.game.NpcPlayer()
  cell := player.SelectTargetCell()
  err := player.PutStoneTo(cell)

  if err != nil {
    return 0, err // FIXME GameResult
  }
  result := f.game.CheckBoard()

  if result == Reached {
    return Lose, nil
  } else if result == Filled {
    return Draw, nil
  } else {
    return 0, nil // FIXME GameResult
  }
}
