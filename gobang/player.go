package gobang

import (
  "encoding/json"
)

func NewGobangPlayer(stone Stone) *GobangPlayer {
  return &GobangPlayer {
    stone: stone,
  }
}

type Player interface {
  Stone() Stone
  PutStoneTo(cell *Cell) error
  MarshalJSON() ([]byte, error)
}

type GobangPlayer struct {
  stone Stone
}

func (player *GobangPlayer) Stone() Stone {
  return player.stone
}

func (player *GobangPlayer) PutStoneTo(cell *Cell) error {
  if !cell.IsEmpty() {
    return NewAlreadyPlacedError(cell)
  }
  player.stone.PutTo(cell)

  return nil
}

func (player *GobangPlayer) MarshalJSON() ([]byte, error) {
  jsonObject := struct {
    Stone Stone `json:"stone"`
  }{
    Stone: player.Stone(),
  }
  return json.Marshal(jsonObject)
}
