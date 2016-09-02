package server

import (
  "testing"
)

func TestGameStartMessage(t *testing.T) {
  raw := []byte("{ \"type\": \"start\", \"body\": { \"stone\": 1 } }")
  msgBody, err := DecodeMessage(raw)

  if err != nil {
    t.Errorf("got %v\n", err)
  }

  switch msg := msgBody.(type) {
  default:
    t.Errorf("got %v\n", msg)
  case *GameStartMessage:
    if msg.Stone != 1 {
      t.Errorf("got %v\nwant %v\n", msg.Stone, 1)
    }
  }
}

func TestSelectCellMessage(t *testing.T) {
  raw := []byte("{ \"type\": \"selectCell\", \"body\": { \"x\": 1, \"y\": 2 } }")
  msgBody, err := DecodeMessage(raw)

  if err != nil {
    t.Errorf("got %v\n", err)
  }

  switch msg := msgBody.(type) {
  default:
    t.Errorf("got %v\n", msg)
  case *SelectCellMessage:
    if msg.X() != 1 {
      t.Errorf("got %v\nwant %v\n", msg.X(), 1)
    }
    if msg.Y() != 2 {
      t.Errorf("got %v\nwant %v\n", msg.X(), 2)
    }
  }
}
