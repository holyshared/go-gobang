package gobang

import (
  "encoding/json"
)

type GameFacilitator interface {
  PlayerPutStoneTo(point *Point) (GameProgressResult, error)
  NpcPlayerPutStone() (GameProgressResult, error)
}

func NewGobang(rule *GameRule, playerStone, npcPlayerStone Stone) *Gobang {
  return &Gobang {
    ctx: NewGameContext(rule, playerStone, npcPlayerStone),
  }
}

type Gobang struct {
  ctx *GameContext
}

func (g *Gobang) PlayerPutStoneTo(point *Point) (GameProgressResult, error) {
  cell, err := g.playerSelectCell(point)

  if err != nil {
    return Retry, err
  }
  result, err := g.playerPutStoneTo(cell)

  if err != nil {
    return Retry, err
  }
  g.changeToNextPlayer()

  return result, nil 
}

func (g *Gobang) NpcPlayerPutStone() (GameProgressResult, error) {
  result, err := g.npcPlayerPutStone()

  if err != nil {
    return Retry, err //FIXME Fatal
  }

  return result, nil 
}

func (g *Gobang) playerSelectCell(point *Point) (*Cell, error) {
  player := g.ctx.GamePlayer()
  cell, err := player.SelectBoardCell(point)

  return cell, err
}

func (g *Gobang) playerPutStoneTo(cell *Cell) (GameProgressResult, error) {
  player := g.ctx.GamePlayer()
  err := player.PutStoneTo(cell)

  if err != nil {
    return Retry, err
  }
  result := g.ctx.CheckBoard()

  if result == Reached {
    return Win, nil
  } else if result == Filled {
    return Draw, nil
  }

  return Next, nil
}

func (g *Gobang) npcPlayerPutStone() (GameProgressResult, error) {
  player := g.ctx.NpcPlayer()
  cell := player.SelectTargetCell()
  err := player.PutStoneTo(cell)

  if err != nil {
    return Retry, err
  }
  result := g.ctx.CheckBoard()

  if result == Reached {
    return Lose, nil
  } else if result == Filled {
    return Draw, nil
  }

  return Next, nil
}

func (g *Gobang) changeToNextPlayer() {
  g.ctx.ChangeToNextPlayer()
}

func (g *Gobang) MarshalJSON() ([]byte, error) {
  return json.Marshal(g.ctx)
}
