package gobang

import (
  "testing"
  "encoding/json"
)

func TestGobang(t *testing.T) {
  var err error
  var result GameProgressResult
  var state []byte

  game := NewGobang(
    DefaultGameRule(),
    Black,
    White,
  )
  result, err = game.PlayerPutStoneTo(NewPoint(5, 1))

  if err != nil {
    t.Errorf("got error %v\n", err)
  }

  if result != Next {
    t.Errorf("got %v\nwant %v", result, Next)
  }

  state, err = json.Marshal(game)

  if err != nil {
    t.Errorf("got error %v\n", err)
  }

  if state == nil {
    t.Errorf("got %v\n", state)
  }
}
