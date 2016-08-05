package gobang

type GameFacilitator struct {
  game *GameContext
}

func (f *GameFacilitator) PlayerPutStone(point *Point) (GameResult, error) {
  player := f.game.GamePlayer()
  result, err := player.PutStone(point)

  if err != nil {
    return 0, err // FIXME GameResult
  }

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
  result, err := player.PutStone(cell.Point)

  if err != nil {
    return 0, err // FIXME GameResult
  }

  if result == Reached {
    return Lose, nil
  } else if result == Filled {
    return Draw, nil
  } else {
    return 0, nil // FIXME GameResult
  }
}
