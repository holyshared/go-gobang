package gobang

func NewGameFacilitator(ctx *GameContext) *GameFacilitator {
  return &GameFacilitator {
    ctx: ctx,
  }
}

type GameFacilitator struct {
  ctx *GameContext
}

func (f *GameFacilitator) playerSelectCell(point *Point) (*Cell, error) {
  player := f.ctx.GamePlayer()
  cell, err := player.SelectBoardCell(point)

  return cell, err
}

func (f *GameFacilitator) PlayerPutStoneTo(point *Point) (GameProgressResult, error) {

//  player := f.ctx.GamePlayer()
  cell, err := f.playerSelectCell(point)

  if err != nil {
    return Retry, err
  }
  result, err := f.playerPutStoneTo(cell)

  if err != nil {
    return Retry, err
  }
  f.changeToNextPlayer()

  return result, err 

//  return f.playerPutStoneTo(cell)
  /*
  player := f.ctx.GamePlayer()
  err := player.PutStoneTo(cell)

  if err != nil {
    return Retry, err
  }
  result := f.ctx.CheckBoard()

  if result == Reached {
    return Win, nil
  } else if result == Filled {
    return Draw, nil
  } else {
    return Next, nil // FIXME GameResult
  }*/
}


func (f *GameFacilitator) playerPutStoneTo(cell *Cell) (GameProgressResult, error) {
  player := f.ctx.GamePlayer()
  err := player.PutStoneTo(cell)

  if err != nil {
    return Retry, err
  }
  result := f.ctx.CheckBoard()

  if result == Reached {
    return Win, nil
  } else if result == Filled {
    return Draw, nil
  }
//  } else {
  return Next, nil // FIXME GameResult
  //}
}





func (f *GameFacilitator) changeToNextPlayer() {
  f.ctx.ChangeToNextPlayer()
}

func (f *GameFacilitator) NpcPlayerPutStone() (GameProgressResult, error) {
  result, err := f.npcPlayerPutStone()

  if err != nil {
    return Retry, err //FIXME Fatal
  }

  return result, nil 
//  result := f.ctx.CheckBoard()

//  if result == Reached {
  //  return Lose, nil
  //} else if result == Filled {
  //  return Draw, nil
//  } else {
//    return Next, nil // FIXME GameResult
  //}
}

func (f *GameFacilitator) npcPlayerPutStone() (GameProgressResult, error) {
  player := f.ctx.NpcPlayer()
  cell := player.SelectTargetCell()
  err := player.PutStoneTo(cell)

  if err != nil {
    return Retry, err
  }
  result := f.ctx.CheckBoard()

  if result == Reached {
    return Lose, nil
  } else if result == Filled {
    return Draw, nil
  }

  return Next, nil
}
